package fuzz_greatape_activitypub

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/reiver/greatape/app/activitypub"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                activitypub.UnmarshalActor(data)
                return 0

            case 1:
                activitypub.UnmarshalOrderedCollection(data)
                return 0

            case 2:
                activitypub.UnmarshalFollowers(data)
                return 0

            case 3:
                activitypub.UnmarshalNote(data)
                return 0

            case 4:
                activitypub.UnmarshalObject(data)
                return 0

            case 5:
                activitypub.UnmarshalOutbox(data)
                return 0

            case 6:
                from, _ := fuzzConsumer.GetString()
                to, _ := fuzzConsumer.GetString()
                content, _ := fuzzConsumer.GetString()

                activitypub.NewNote(from, to, content)
                return 0

            case 7:
                var testNote activitypub.Note
                fuzzConsumer.GenerateStruct(&testNote)

                username, _ := fuzzConsumer.GetString()
                publicUrl, _ := fuzzConsumer.GetString()
                uniqueIdentifier, _ := fuzzConsumer.GetString()

                testNote.Wrap(username, publicUrl, uniqueIdentifier)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}