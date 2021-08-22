package modle

type List_Song struct {

	Id int64 `json:"id" gorm:"primarykey; autoIncrement:true"`
	Name_Song string `json:"name_song"`
	Path string `json:"path"`
	Song_id string `json:"song_id"`

}

