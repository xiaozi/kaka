package main

func main() {
	config := NewConfig()
	browser := NewMacBrowser(config)
	browser.Snapshot("http://tool.lu/", "tool.jpg")
	storage := NewStorage(config)
	storage.put("remote/path.txt", "testfile.txt")
}
