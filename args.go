package main

import "flag"

func ParseArgs() (postId int) {
	flag.IntVar(&postId, "id", 0, "Target post ID on esa.io")
	flag.Parse()

	return postId
}
