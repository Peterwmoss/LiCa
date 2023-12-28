package database

import (
	"context"
	"log"

	"github.com/uptrace/bun"
)

func Seed(db *bun.DB, ctx context.Context) error {
	err := seedCategories(db, ctx)
	if err != nil {
		return err
	}

	err = seedItems(db, ctx)
	if err != nil {
		return err
	}

	return nil
}

func seedCategories(db *bun.DB, ctx context.Context) error {
	categories := []Category{
		{Name: "mejeri"},
		{Name: "pålæg"},
		{Name: "kød/fisk/fjerkræ"},
		{Name: "kolonial"},
		{Name: "brød"},
		{Name: "husholdning"},
		{Name: "personlig pleje"},
		{Name: "snacks"},
		{Name: "slik/chokolade"},
		{Name: "sodavand/øl"},
		{Name: "vin/spiritus"},
		{Name: "bolig/fritid"},
		{Name: "frost"},
		{Name: "frugt/grønt"},
		{Name: "andet"},
	}

	return insertNoDuplicates(db, categories, ctx)
}

func seedItems(db *bun.DB, ctx context.Context) error {
	items := []Item{
		{Name: "mælk", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "smør", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "smørbart", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "skyr", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "creme fraiche", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "æg", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "piskefløde", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "madlavningsfløde", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "ost", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "parmesan", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "cheddar", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "revet cheddar", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "mozzerella", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "frisk mozzerella", CategoryId: findCategoryByName("mejeri", db, ctx).Id},
		{Name: "revet mozzerella", CategoryId: findCategoryByName("mejeri", db, ctx).Id},

		{Name: "skivepålæg", CategoryId: findCategoryByName("pålæg", db, ctx).Id},
		{Name: "klappålæg", CategoryId: findCategoryByName("pålæg", db, ctx).Id},
		{Name: "laksesalat", CategoryId: findCategoryByName("pålæg", db, ctx).Id},
		{Name: "røget laks", CategoryId: findCategoryByName("pålæg", db, ctx).Id},

		{Name: "laks", CategoryId: findCategoryByName("kød/fisk/fjerkræ", db, ctx).Id},
		{Name: "hakket oksekød", CategoryId: findCategoryByName("kød/fisk/fjerkræ", db, ctx).Id},
		{Name: "kylling inderfilet", CategoryId: findCategoryByName("kød/fisk/fjerkræ", db, ctx).Id},
		{Name: "andebryst", CategoryId: findCategoryByName("kød/fisk/fjerkræ", db, ctx).Id},
		{Name: "hakket svinekød", CategoryId: findCategoryByName("kød/fisk/fjerkræ", db, ctx).Id},
		{Name: "svinemørbrad", CategoryId: findCategoryByName("kød/fisk/fjerkræ", db, ctx).Id},
		{Name: "oksemørbrad", CategoryId: findCategoryByName("kød/fisk/fjerkræ", db, ctx).Id},
		{Name: "torsk", CategoryId: findCategoryByName("kød/fisk/fjerkræ", db, ctx).Id},

		{Name: "ris", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "pasta", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "soya", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "majs", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "tun i olie", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "lagereddike", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "ketchup", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "mayonnaise", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "remoulade", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "cruesli", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "nudler", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "risnudler", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "ægnudler", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "tortilla", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "taco", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "artiskokhjerter", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "hakkede tomater", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "pesto", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "olivenolie", CategoryId: findCategoryByName("kolonial", db, ctx).Id},
		{Name: "bearnaise", CategoryId: findCategoryByName("kolonial", db, ctx).Id},

		{Name: "boller", CategoryId: findCategoryByName("brød", db, ctx).Id},
		{Name: "rugbrød", CategoryId: findCategoryByName("brød", db, ctx).Id},
		{Name: "burgerboller", CategoryId: findCategoryByName("brød", db, ctx).Id},

		{Name: "toiletpapir", CategoryId: findCategoryByName("husholdning", db, ctx).Id},
		{Name: "køkkenrulle", CategoryId: findCategoryByName("husholdning", db, ctx).Id},
		{Name: "opvasketabs", CategoryId: findCategoryByName("husholdning", db, ctx).Id},
		{Name: "vakemiddel farve", CategoryId: findCategoryByName("husholdning", db, ctx).Id},
		{Name: "vakemiddel hvid", CategoryId: findCategoryByName("husholdning", db, ctx).Id},
		{Name: "opvaskesæbe", CategoryId: findCategoryByName("husholdning", db, ctx).Id},
		{Name: "eddikesyre", CategoryId: findCategoryByName("husholdning", db, ctx).Id},
		{Name: "afløbsrens", CategoryId: findCategoryByName("husholdning", db, ctx).Id},
		{Name: "toiletrens", CategoryId: findCategoryByName("husholdning", db, ctx).Id},
		{Name: "fryseposer", CategoryId: findCategoryByName("husholdning", db, ctx).Id},

		{Name: "tandpasta", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "ansigtscreme", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "håndcreme", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "creme", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "læbepommade", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "mundskyl", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "natbind", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "trusseindlæg", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "alm. bind", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "tamponer superplus", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "tamponer super", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "tamponer regular", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "tamponer light", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},
		{Name: "tandtråd", CategoryId: findCategoryByName("personlig pleje", db, ctx).Id},

		{Name: "chips", CategoryId: findCategoryByName("snacks", db, ctx).Id},
		{Name: "popcorn", CategoryId: findCategoryByName("snacks", db, ctx).Id},
		{Name: "dip", CategoryId: findCategoryByName("snacks", db, ctx).Id},
		{Name: "rugbrødschips", CategoryId: findCategoryByName("snacks", db, ctx).Id},
		{Name: "snacks", CategoryId: findCategoryByName("snacks", db, ctx).Id},

		{Name: "click mix", CategoryId: findCategoryByName("slik/chokolade", db, ctx).Id},
		{Name: "slik", CategoryId: findCategoryByName("slik/chokolade", db, ctx).Id},
		{Name: "chokolade", CategoryId: findCategoryByName("slik/chokolade", db, ctx).Id},

		{Name: "sodavand", CategoryId: findCategoryByName("sodavand/øl", db, ctx).Id},
		{Name: "øl", CategoryId: findCategoryByName("sodavand/øl", db, ctx).Id},
		{Name: "cola", CategoryId: findCategoryByName("sodavand/øl", db, ctx).Id},
		{Name: "redbull", CategoryId: findCategoryByName("sodavand/øl", db, ctx).Id},
		{Name: "ginger beer", CategoryId: findCategoryByName("sodavand/øl", db, ctx).Id},

		{Name: "rødvin", CategoryId: findCategoryByName("vin/spiritus", db, ctx).Id},
		{Name: "hvidvin", CategoryId: findCategoryByName("vin/spiritus", db, ctx).Id},
		{Name: "vin", CategoryId: findCategoryByName("vin/spiritus", db, ctx).Id},
		{Name: "alkohol", CategoryId: findCategoryByName("vin/spiritus", db, ctx).Id},
		{Name: "gin", CategoryId: findCategoryByName("vin/spiritus", db, ctx).Id},
		{Name: "vodka", CategoryId: findCategoryByName("vin/spiritus", db, ctx).Id},
		{Name: "lys rom", CategoryId: findCategoryByName("vin/spiritus", db, ctx).Id},
		{Name: "mørk rom", CategoryId: findCategoryByName("vin/spiritus", db, ctx).Id},
		{Name: "rom", CategoryId: findCategoryByName("vin/spiritus", db, ctx).Id},

		{Name: "aaa batterier", CategoryId: findCategoryByName("bolig/fritid", db, ctx).Id},
		{Name: "aa batterier", CategoryId: findCategoryByName("bolig/fritid", db, ctx).Id},
		{Name: "lim", CategoryId: findCategoryByName("bolig/fritid", db, ctx).Id},
		{Name: "saks", CategoryId: findCategoryByName("bolig/fritid", db, ctx).Id},
		{Name: "printerpapir", CategoryId: findCategoryByName("bolig/fritid", db, ctx).Id},
		{Name: "papir", CategoryId: findCategoryByName("bolig/fritid", db, ctx).Id},

		{Name: "frosne bær", CategoryId: findCategoryByName("frost", db, ctx).Id},
		{Name: "is", CategoryId: findCategoryByName("frost", db, ctx).Id},
		{Name: "thai boks", CategoryId: findCategoryByName("frost", db, ctx).Id},
		{Name: "frysepizza", CategoryId: findCategoryByName("frost", db, ctx).Id},

		{Name: "bær", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "agurk", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "gulerod", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "ananas", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "tomater", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "bananer", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "spinat", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "salat", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "icebergsalat", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "hjertesalat", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "ruccola", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "appelsiner", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "æbler", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "passionsfrugt", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "spidskål", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "løg", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "kartofler", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "rødløg", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},
		{Name: "forårsløg", CategoryId: findCategoryByName("frugt/grønt", db, ctx).Id},

		{Name: "kulsyrepatron", CategoryId: findCategoryByName("andet", db, ctx).Id},
	}

	return insertNoDuplicates(db, items, ctx)
}

func findCategoryByName(name string, db *bun.DB, ctx context.Context) *Category {
	category := &Category{}
	err := db.NewSelect().
		Model(category).
		Where("name = ?", name).
		Limit(1).
		Scan(ctx)

	if err != nil {
		log.Fatal(err)
	}

	if category == nil {
		log.Fatal("Could not find category with name. Failed to seed")
	}

	return category
}

func insertNoDuplicates[T interface{}](db *bun.DB, models []T, ctx context.Context) error {
	for _, model := range models {
		_, err := db.
			NewInsert().
			Model(&model).
			On("CONFLICT DO NOTHING").
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
