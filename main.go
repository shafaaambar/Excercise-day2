package main

import (
	"oop/blog"
	"oop/product"
	"oop/profile"
	"oop/store"
)

func main() {
	var name = "Ambar"
	profile := profile.User{}

	profile.SetProfile(name, "kamu", "dimana")
	profile.GetProfile()

	product := product.NEW("Aqua", 100)
	product.GetProduct()

	blog := blog.New("Tutorial Golang", "OOP inheritance", "Ambar")
	blog.GetBlog()

	branchStore := store.BranchStore{StoreName: "cabang", OwnerBranch: "Ambar"}
	store := store.Store{BranchStore: branchStore, StoreName: "cabang", Owner: "Ambar"}

	store.GetBranchStore()
}
