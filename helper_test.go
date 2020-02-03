package main

import (
	"reflect"
	"sort"
	"testing"
)

func IsEqual(a1 []string, a2 []string) bool {
	sort.Strings(a1)
	sort.Strings(a2)
	if len(a1) == len(a2) {
		for i, v := range a1 {
			if v != a2[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func TestGetUserIDByUsername(t *testing.T) {
	userName := [10]string{"Rob", "icemanphotos", "Trump", ".....", "AMS", "Eric Fischer", "chetangarg365", "oprisco", "pavonne", "bolandrotor"}
	userId := [10]string{"35468141087@N01", "42957889@N05", "61607373@N00", "87464910@N00", "43564525@N06", "24431382@N03", "186712313@N06", "35116247@N03", "18114173@N07", "79649916@N00"}
	for i := 0; i < 10; i++ {
		ui, err := GetUserIDByUsername(userName[i])
		if err != nil {
			t.Errorf("%s", err)
			t.Errorf("want: %s got: %s", userId[i], ui)
		} else {
			if ui != userId[i] {
				t.Errorf("want: %s got: %s", userId[i], ui)
				t.Errorf("Username: %s", userName[i])
			}
		}
		//GetAlbumsFromUserID("35468141087@N01")
	}
}

func TestGetAlbumsFromUserID(t *testing.T) {
	userAlbum := [][]string{
		{"Flickr Top 25", "Best and Most", "Lux* South Ari Atoll", "Gili Lankanfushi Resort", "Bandos Resort", "Paradise Island Resort and Spa", "Explored", "Getty Images [for sale]", "Project: Waterdrops", "ButterFlies", "Project: Smoke", "Product Photography", "Flowers", "Project: Kitchen Staff", "Christmas"},
		{"18mm T*", "All Saints", "B&W", "Basket Case", "Corpus Christi", "Day by the water", "Getty Want", "Ionosphere", "KLOSZart.Collective", "Michael's swing", "Order of dada", "Pano Pano Rama Rama", "Praha", "Travel", "U Ani", "Yard", "Zenit", "book", "circles,lines,squares", "egos", "elevationz", "few birds", "human vs indoors", "old house", "the incredible herkules", "warmth from the bottom of my guts", "Česká republika", "żenia"},
		{"Bittern", "Owls", "Birds of Prey Workshop", "Kingfishers", "Mammals", "Raptors", "Seabirds", "Heart of Midlothian", "Plants", "Corvidae", "Edinburgh", "Amphibians & Reptiles", "Land/Seascapes", "Insects", "Birds", "Waders", "Ducks/Geese/Swans", "Seaside", "Woodpecker"},
		{"100 Strangers", "Under the Sea", "The Thirty Dollar Project", "Leaves", "Interestingness Set", "Animals", "Buildings & Structures", "Birds", "Black & White", "bugs & Insects", "Cactus", "Experimenting with Light", "Fruit & Vegetables", "Flowers", "Instruments", "Nature", "People", "sunrises & Sunsets", "Water", "Random"}}
	userId := [4]string{"42957889@N05", "79649916@N00", "16177003@N03", "18114173@N07"}
	for i := 0; i < 4; i++ {
		album, err := GetAlbumsFromUserID((userId[i]))
		if err != nil {
			t.Errorf("%s", err)
			t.Errorf("want: %v got: %v", userAlbum[i], album)
			t.Errorf("UserId: %s", userId[i])
		} else {
			if !IsEqual(album, userAlbum[i]) {
				t.Errorf("want: %v got: %v", userAlbum[i], album)
				t.Errorf("UserId: %s", userId[i])
			}
		}
	}
}

//16177003@N03
func TestGetPhotosFromAlbum(t *testing.T) {
	userIds := []string{"79649916@N00", "79649916@N00", "16177003@N03"}
	photoSetIds := []string{"72157680297779204", "72157689249575040", "72157629858997866"}
	var photoTests = [3]map[string]string{}
	photoTests[0] = map[string]string{
		"33645529354": "Praha.jpg", "33645647574": "Praha.jpg", "33651289754": "Praha.jpg", "33677799603": "Praha.jpg", "33678269633": "Praha.jpg", "34102641220": "Praha.jpg", "34102828170": "Praha.jpg", "34103163720": "Praha.jpg", "34328112862": "Praha.jpg", "34357573721": "Praha.jpg", "34447069196": "Praha.jpg", "34447101266": "Praha.jpg", "34452173006": "Praha.jpg", "34487897055": "Praha.jpg",
	}
	photoTests[1] = map[string]string{
		"1410924473": "play ball with crazy dandelions.jpg", "29397647236": "Court.jpg", "34035214841": "bd.jpg", "36955369502": "Highlander's private yard.jpg", "40165414815": "Forgotten games.jpg", "40611995495": "Well played.jpg",
	}
	photoTests[2] = map[string]string{
		"7240082798": "Scottish Gallery of Moden Art.jpg", "7250450622": "Scottish Cup Final 2012.jpg",
	}
	for i := 0; i < 3; i++ {
		photos, err := GetPhotosFromAlbum(userIds[i], photoSetIds[i])
		if err != nil {
			t.Errorf("%s", err)
			t.Errorf("want: %v \ngot: %v", photoTests[i], photos)
			t.Errorf("UserId: %s  PhotoSetId:  %s", userIds[i], photoSetIds[i])
		}
		if !reflect.DeepEqual(photos, photoTests[i]) {
			t.Errorf("want: %v \ngot: %v", photoTests[i], photos)
			t.Errorf("UserId: %s  PhotoSetId:  %s", userIds[i], photoSetIds[i])
		}
	}
}

func TestGetPhotoSizes(t *testing.T) {
	photoIds := []string{"40165414815", "29397647236", "36955369502"}
	var photoTests = [3]map[string]string{}
	photoTests[0] = map[string]string{
		"https://live.staticflickr.com/809/40165414815_8a85c08311_n.jpg": "Small 320", "https://live.staticflickr.com/809/40165414815_8a85c08311_w.jpg": "Small 400", "https://live.staticflickr.com/809/40165414815_98c743f46b_5k.jpg": "X-Large 5K", "https://live.staticflickr.com/809/40165414815_8a85c08311.jpg": "Medium", "https://live.staticflickr.com/809/40165414815_8a85c08311_c.jpg": "Medium 800", "https://live.staticflickr.com/809/40165414815_8a85c08311_s.jpg": "Square", "https://live.staticflickr.com/809/40165414815_8a85c08311_m.jpg": "Small", "https://live.staticflickr.com/809/40165414815_3a9ab7de41_h.jpg": "Large 1600", "https://live.staticflickr.com/809/40165414815_2cd156c46f_k.jpg": "Large 2048", "https://live.staticflickr.com/809/40165414815_926d3047af_3k.jpg": "X-Large 3K", "https://live.staticflickr.com/809/40165414815_8a85c08311_q.jpg": "Large Square", "https://live.staticflickr.com/809/40165414815_8a85c08311_t.jpg": "Thumbnail", "https://live.staticflickr.com/809/40165414815_8a85c08311_z.jpg": "Medium 640", "https://live.staticflickr.com/809/40165414815_8a85c08311_b.jpg": "Large", "https://live.staticflickr.com/809/40165414815_bbd8a2cea3_4k.jpg": "X-Large 4K", "https://live.staticflickr.com/809/40165414815_e446ec3754_6k.jpg": "X-Large 6K",
	}
	photoTests[1] = map[string]string{
		"https://live.staticflickr.com/8300/29397647236_1921449bb8_q.jpg": "Large Square", "https://live.staticflickr.com/8300/29397647236_544a94c4e1_3k.jpg": "X-Large 3K", "https://live.staticflickr.com/8300/29397647236_32b3041f93_4k.jpg": "X-Large 4K", "https://live.staticflickr.com/8300/29397647236_1921449bb8_b.jpg": "Large", "https://live.staticflickr.com/8300/29397647236_3c8745cf25_h.jpg": "Large 1600", "https://live.staticflickr.com/8300/29397647236_1921449bb8_m.jpg": "Small", "https://live.staticflickr.com/8300/29397647236_1921449bb8_n.jpg": "Small 320", "https://live.staticflickr.com/8300/29397647236_1921449bb8_w.jpg": "Small 400", "https://live.staticflickr.com/8300/29397647236_1921449bb8.jpg": "Medium", "https://live.staticflickr.com/8300/29397647236_1921449bb8_z.jpg": "Medium 640", "https://live.staticflickr.com/8300/29397647236_1921449bb8_c.jpg": "Medium 800", "https://live.staticflickr.com/8300/29397647236_c0cb2874aa_6k.jpg": "X-Large 6K", "https://live.staticflickr.com/8300/29397647236_1921449bb8_s.jpg": "Square", "https://live.staticflickr.com/8300/29397647236_1921449bb8_t.jpg": "Thumbnail", "https://live.staticflickr.com/8300/29397647236_e219d9ea0c_k.jpg": "Large 2048", "https://live.staticflickr.com/8300/29397647236_6effa00f0b_5k.jpg": "X-Large 5K",
	}
	photoTests[2] = map[string]string{
		"https://live.staticflickr.com/4423/36955369502_b73c90cff9_5k.jpg": "X-Large 5K", "https://live.staticflickr.com/4423/36955369502_854628bcb6_w.jpg": "Small 400", "https://live.staticflickr.com/4423/36955369502_854628bcb6_z.jpg": "Medium 640", "https://live.staticflickr.com/4423/36955369502_854628bcb6_c.jpg": "Medium 800", "https://live.staticflickr.com/4423/36955369502_854628bcb6_b.jpg": "Large", "https://live.staticflickr.com/4423/36955369502_854628bcb6_s.jpg": "Square", "https://live.staticflickr.com/4423/36955369502_854628bcb6_m.jpg": "Small", "https://live.staticflickr.com/4423/36955369502_f26e123b69_3k.jpg": "X-Large 3K", "https://live.staticflickr.com/4423/36955369502_854628bcb6_t.jpg": "Thumbnail", "https://live.staticflickr.com/4423/36955369502_4316c3aee4_k.jpg": "Large 2048", "https://live.staticflickr.com/4423/36955369502_f35ea46258_4k.jpg": "X-Large 4K", "https://live.staticflickr.com/4423/36955369502_854628bcb6_q.jpg": "Large Square", "https://live.staticflickr.com/4423/36955369502_854628bcb6_n.jpg": "Small 320", "https://live.staticflickr.com/4423/36955369502_854628bcb6.jpg": "Medium", "https://live.staticflickr.com/4423/36955369502_4e3fe85aaf_h.jpg": "Large 1600",
	}
	for i := 0; i < 3; i++ {
		photos, err := GetPhotoSizes(photoIds[i])
		if err != nil {
			t.Errorf("%s", err)
			t.Errorf("want: %v \ngot: %v", photoTests[i], photos)
			t.Errorf("PhotoId:  %s", photoIds[i])
		}
		if !reflect.DeepEqual(photos, photoTests[i]) {
			t.Errorf("want: %v \ngot: %v", photoTests[i], photos)
			t.Errorf("PhotoId:  %s", photoIds[i])
		}
	}
}
