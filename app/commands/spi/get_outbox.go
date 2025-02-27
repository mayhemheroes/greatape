package spi

import (
	"time"

	"github.com/reiver/greatape/app/activitypub"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func GetOutbox(x IDispatcher, username string) (IGetOutboxResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Username() == username
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_FOUND)
	identity := identities.First()

	actor := x.Format("%s/u/%s", x.PublicUrl(), identity.Username())

	activities := x.FilterActivityPubOutgoingActivities(func(activity IActivityPubOutgoingActivity) bool {
		return activity.From() == actor && activity.To() == ACTIVITY_PUB_PUBLIC
	})

	var orderedItems ActivityPubActivities
	activities.ForEach(func(outgoingActivity IActivityPubOutgoingActivity) {
		published := time.Unix(0, outgoingActivity.Timestamp()).Format("2006-01-02T15:04:05Z")

		note := activitypub.NewPublicNote(actor, outgoingActivity.Content())
		noteActivity := note.Wrap(username, x.PublicUrl(), outgoingActivity.UniqueIdentifier())

		object, _ := x.NewActivityPubObject()
		object.SetContext(ACTIVITY_STREAMS)
		object.SetType(ACTIVITY_PUB_NOTE)
		object.SetId(note.Id)
		object.SetContent(note.Content)

		activity, _ := x.NewActivityPubActivity()
		activity.SetContext(ACTIVITY_STREAMS)
		activity.SetType(ACTIVITY_PUB_CREATE)
		activity.SetId(x.Format("%s/posts/%s", actor, outgoingActivity.UniqueIdentifier()))
		activity.SetActor(actor)
		activity.SetTo(noteActivity.To.([]string))
		activity.SetPublished(published)
		activity.SetObject(object)

		orderedItems = append(orderedItems, activity)
	})

	return x.NewGetOutboxResult(
		ACTIVITY_STREAMS,                // context
		x.Format("%s/outbox", actor),    // id
		ACTIVITY_PUB_ORDERED_COLLECTION, // type
		int32(len(orderedItems)),        // totalItems
		orderedItems,                    // orderedItems
		"",                              // first
	), nil
}
