package model

type Buku struct {
	ID           int64  `json:"id"`
	Kategori_ID  int64  `json:"kategori_id"`
	Judul_Buku   string `json:"judul_buku"`
	Stok         int64  `json:"stok"`
	Penulis      string `json:"penulis"`
	Penerbit     string `json:"penerbit"`
	Uraian       string `json:"uraian"`
	Tahun_Terbit int64  `json:"tahun_terbit"`
	ISBN         string `json:"isbn"`
	Sumber       string `json:"sumber"`
	Kode_Tempat  string `json:"kode_tempat"`
}
