package tags

import (
	"github.com/lleoserrano/fiber_mongo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sort"
)

func FindOrCreate(name string) (doc Tag, err error) {
	filter := bson.M{"name": name}

	err = db.FindOne("tags", filter, &doc)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}

	if doc.Name != "" {
		return
	}

	doc.Name = name

	id, err := db.Insert("tags", doc)
	if err != nil {
		return
	}

	doc.ID = id

	return
}

func AddTask(taskID string, names []string) error {
	for _, name := range names {
		tag, err := FindOrCreate(name)
		if err != nil {
			return err
		}
		i := sort.SearchStrings(tag.Tasks, taskID)

		if i < len(tag.Tasks) && tag.Tasks[i] == taskID {
			continue
		}

		tag.Tasks = append(tag.Tasks, taskID)
		sort.Strings(tag.Tasks)
		result := new(Tag)
		err = db.UpdateByID("tags", tag.ID.Hex(), tag, &result)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveTask(taskID string, names ...string) error {
	filter := bson.M{"tasks": taskID}

	if len(names) > 0 {
		filter["name"] = bson.M{"$in": names}
	}

	var documents []Tag

	err := db.FindAll("tags", filter, &documents)
	if err != nil {
		return err
	}

	var result Tag

	for _, document := range documents {
		i := sort.SearchStrings(document.Tasks, taskID)

		document.Tasks = append(document.Tasks[:i], document.Tasks[i+1:]...)
		err := db.UpdateByID("tags", document.ID.Hex(), document, &result)
		if err != nil {
			return err
		}
	}
	return nil
}
