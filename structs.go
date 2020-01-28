package main

type UserStruct struct {
	User struct {
		ID       string `json:"id"`
		Nsid     string `json:"nsid"`
		Username struct {
			Content string `json:"_content"`
		} `json:"username"`
	} `json:"user"`
	Stat string `json:"stat"`
}

type Album struct {
	Photosets struct {
		Page     int    `json:"page"`
		Pages    int    `json:"pages"`
		Perpage  string `json:"perpage"`
		Total    string `json:"total"`
		Photoset []struct {
			ID            string `json:"id"`
			Owner         string `json:"owner"`
			Username      string `json:"username"`
			Primary       string `json:"primary"`
			Secret        string `json:"secret"`
			Server        string `json:"server"`
			Farm          int    `json:"farm"`
			CountViews    string `json:"count_views"`
			CountComments int    `json:"count_comments"`
			CountPhotos   int    `json:"count_photos"`
			CountVideos   int    `json:"count_videos"`
			Title         struct {
				Content string `json:"_content"`
			} `json:"title"`
			Description struct {
				Content string `json:"_content"`
			} `json:"description"`
			CanComment          int    `json:"can_comment"`
			DateCreate          string `json:"date_create"`
			DateUpdate          string `json:"date_update"`
			Photos              int    `json:"photos"`
			Videos              int    `json:"videos"`
			VisibilityCanSeeSet int    `json:"visibility_can_see_set"`
			NeedsInterstitial   int    `json:"needs_interstitial"`
		} `json:"photoset"`
	} `json:"photosets"`
	Stat string `json:"stat"`
}

type AlbumPhotos struct {
	Photoset struct {
		ID        string `json:"id"`
		Primary   string `json:"primary"`
		Owner     string `json:"owner"`
		Ownername string `json:"ownername"`
		Photo     []struct {
			ID        string `json:"id"`
			Secret    string `json:"secret"`
			Server    string `json:"server"`
			Farm      int    `json:"farm"`
			Title     string `json:"title"`
			Isprimary int    `json:"isprimary"`
			Ispublic  int    `json:"ispublic"`
			Isfriend  int    `json:"isfriend"`
			Isfamily  int    `json:"isfamily"`
		} `json:"photo"`
		Page    int    `json:"page"`
		PerPage string `json:"per_page"`
		Perpage string `json:"perpage"`
		Pages   int    `json:"pages"`
		Title   string `json:"title"`
		Total   int    `json:"total"`
	} `json:"photoset"`
	Stat string `json:"stat"`
}

type PhotoSizes struct {
	Sizes struct {
		Canblog     int `json:"canblog"`
		Canprint    int `json:"canprint"`
		Candownload int `json:"candownload"`
		Size        []struct {
			Label  string `json:"label"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Source string `json:"source"`
			URL    string `json:"url"`
			Media  string `json:"media"`
		} `json:"size"`
	} `json:"sizes"`
	Stat string `json:"stat"`
}
