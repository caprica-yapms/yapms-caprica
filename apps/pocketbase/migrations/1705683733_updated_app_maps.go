package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("3grvxt9a57oyktx")
		if err != nil {
			return err
		}

		// update
		edit_set := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jcsn0tdz",
			"name": "set",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_set)
		collection.Schema.AddField(edit_set)

		// update
		edit_country := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "lvbcfavi",
			"name": "country",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_country)
		collection.Schema.AddField(edit_country)

		// update
		edit_type := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "11zqc9lf",
			"name": "type",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_type)
		collection.Schema.AddField(edit_type)

		// update
		edit_title := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "a4hadbgl",
			"name": "title",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_title)
		collection.Schema.AddField(edit_title)

		// update
		edit_file := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "a8p7yqzz",
			"name": "file",
			"type": "file",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), edit_file)
		collection.Schema.AddField(edit_file)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("3grvxt9a57oyktx")
		if err != nil {
			return err
		}

		// update
		edit_set := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jcsn0tdz",
			"name": "set",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_set)
		collection.Schema.AddField(edit_set)

		// update
		edit_country := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "lvbcfavi",
			"name": "country",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_country)
		collection.Schema.AddField(edit_country)

		// update
		edit_type := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "11zqc9lf",
			"name": "type",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_type)
		collection.Schema.AddField(edit_type)

		// update
		edit_title := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "a4hadbgl",
			"name": "title",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_title)
		collection.Schema.AddField(edit_title)

		// update
		edit_file := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "a8p7yqzz",
			"name": "file",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), edit_file)
		collection.Schema.AddField(edit_file)

		return dao.SaveCollection(collection)
	})
}
