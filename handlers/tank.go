package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	gomongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"killifish/docs"
	"killifish/mongo"
)

var size = []D{{{
	"$group",
	M{
		"_id":    "$size",
		"tanks":  M{"$sum": 1},
		"fishes": M{"$sum": "$amount"},
	},
}}}

var sexual = []D{{{
	"$group",
	M{
		"_id":    "$sexual",
		"tanks":  M{"$sum": 1},
		"fishes": M{"$sum": "$amount"},
	},
}}}

var species = []D{{{
	"$group",
	M{
		"_id":    "$species",
		"tanks":  M{"$sum": 1},
		"fishes": M{"$sum": "$amount"},
	},
}}}

var genotype = []D{{{
	"$group",
	M{
		"_id":    "$genotype",
		"tanks":  M{"$sum": 1},
		"fishes": M{"$sum": "$amount"},
	},
}}}

var birthday = func() []D {
	now := time.Now()
	layout := "2006-01-02"

	Past := func(t time.Time, day int) string {
		d := time.Duration(-24 * day)
		return t.Add(time.Hour * d).Format(layout)
	}
	boundaries := make([]string, 0)

	boundaries = append(boundaries, Past(now, 270))
	for i := 16; i >= 0; i-- {
		boundaries = append(boundaries, Past(now, i*7))
	}

	return []D{{{
		"$bucket",
		M{
			"groupBy":    "$birthday",
			"boundaries": boundaries,
			"default":    "unset",
			"output": M{
				"tanks":  M{"$sum": 1},
				"fishes": M{"$sum": "$amount"},
			},
		},
	}}}
}

func StatTanks(c *fiber.Ctx) error {
	ctx := c.Context()
	name := c.GetRespHeader("User-Name")

	match := D{{"$match", M{"owner": name}}}
	facet := D{{
		"$facet",
		M{
			"size":     size,
			"sexual":   sexual,
			"species":  species,
			"genotype": genotype,
			"birthday": birthday(),
		},
	}}
	pipeline := gomongo.Pipeline{match, facet}

	cursor, e1 := mongo.Tanks.Aggregate(ctx, pipeline)
	if e1 != nil {
		fmt.Print("error1:")
		fmt.Println(e1)
		return e1
	}

	stats := make([]docs.Facets, 0)
	e2 := cursor.All(c.Context(), &stats)
	if e2 != nil {
		fmt.Print("error2:")
		fmt.Println(e2)
		return e2
	}

	return c.JSON(stats[0])
}

func StatAllTanks(c *fiber.Ctx) error {
	ctx := c.Context()

	facet := D{{
		"$facet",
		M{
			"species":  species,
			"size":     size,
			"genotype": genotype,
			"birthday": birthday(),
		},
	}}
	pipeline := gomongo.Pipeline{facet}

	cursor, e1 := mongo.Tanks.Aggregate(ctx, pipeline)
	if e1 != nil {
		fmt.Print("error1:")
		fmt.Println(e1)
		return e1
	}

	stats := make([]D, 0)
	e2 := cursor.All(c.Context(), &stats)
	if e2 != nil {
		fmt.Print("error2:")
		fmt.Println(e2)
		return e2
	}

	return c.JSON(stats[0])
}

func GetTank(c *fiber.Ctx) error {
	id := c.Params("id")
	f := M{"_id": id}
	t := docs.Tank{}
	if err := mongo.Tanks.FindOne(c.Context(), f).Decode(&t); err != nil {
		return c.Status(InternalServerError).JSON(gomongo.ErrNoDocuments)
	}

	return c.JSON(t)
}

type Filter struct {
	Owner  *string `json:"owner,omitempty"`
	Size   *int    `json:"size,omitempty"`
	Sexual *string `json:"sexual,omitempty"`
}

type Search struct {
	Filter fiber.Map `json:"filter"`
	Sort   fiber.Map `json:"sort,omitempty"`
}

func GetTanks(c *fiber.Ctx) error {
	ctx := c.Context()
	s := Search{}
	_ = json.Unmarshal(c.Body(), &s)

	f := s.Filter
	//count, _ := mongo.Tanks.CountDocuments(ctx, f)
	o := options.Find().
		SetSort(s.Sort)

	cursor, _ := mongo.Tanks.Find(ctx, f, o)

	tanks := make([]docs.Tank, 0)
	_ = cursor.All(ctx, &tanks)

	//return c.JSON(
	//	fiber.Map{
	//		"count": count,
	//		"tanks": tanks,
	//	},
	//)
	return c.JSON(tanks)
}

func GetAllTanks(c *fiber.Ctx) error {
	ctx := c.Context()
	f := M{}

	cursor, _ := mongo.Tanks.Find(ctx, f)

	tanks := make([]docs.Tank, 0)
	_ = cursor.All(ctx, tanks)

	return c.JSON(tanks)
}

func InsertTank(c *fiber.Ctx) error {
	t := docs.Tank{}
	_ = c.BodyParser(&t)

	r, e := mongo.Tanks.InsertOne(c.Context(), t)

	if e != nil {
		if gomongo.IsDuplicateKeyError(e) {
			return c.JSON("exist")
		} else {
			return c.JSON("server error")
		}
	}

	return c.JSON(r)
}

func InsertTanks(c *fiber.Ctx) error {
	ts := make([]interface{}, 0)
	_ = c.BodyParser(ts)

	r, _ := mongo.Tanks.InsertMany(c.Context(), ts)

	return c.JSON(r)
}

func UpdateTank(c *fiber.Ctx) error {
	id := c.Params("id")
	m := fiber.Map{}

	_ = json.Unmarshal(c.Body(), &m)
	u := M{"$set": m}

	fmt.Println(u)
	r, _ := mongo.Tanks.UpdateByID(c.Context(), id, u)

	return c.JSON(r)
}

func UpdateTanks(c *fiber.Ctx) error {
	var ids []string
	f := M{"_id": M{"$in": ids}}

	m := fiber.Map{}
	e := json.Unmarshal(c.Body(), &m)

	//fmt.Println(string(c.Body()))
	//fmt.Println(m)
	if e != nil {
		fmt.Println(e)
	}

	u := M{"$set": m}
	fmt.Println(u)
	r, _ := mongo.Tanks.UpdateMany(c.Context(), f, u)

	return c.JSON(r)
}

func DeleteTank(c *fiber.Ctx) error {
	id := c.Params("id")
	f := M{"_id": id}

	r, _ := mongo.Tanks.DeleteOne(c.Context(), f)

	return c.JSON(r)
}

func DeleteTanks(c *fiber.Ctx) error {
	ids := make([]string, 0)
	f := M{"_id": M{"$in": ids}}

	r, _ := mongo.Tanks.DeleteMany(c.Context(), f)

	return c.JSON(r)
}
