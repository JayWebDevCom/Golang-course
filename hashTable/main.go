package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// runes.RuneReminder()
	// fmt.Println(hashBucketToAssignWordTo("my word", 5))
	//	loremReading()
	// mobyDick()
	makeHashTable()
}

func makeHashTable() {
	// get sherlockHolmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal("There was an error retrieving", err)
	}

	// scan the returned page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set the split finction for the scanning operation
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts in each bucket
	buckets := make([]int32, 12)
	// create a map hold counts in each bucket
	bucketsMap := make(map[int]int, 12)
	// make hash-letter-buckets - num of words starting with that letter
	hashLetterBuckets := make([]int, 300)

	// loop over the words
	for scanner.Scan() {
		word := scanner.Text()

		hashBucketToAssignWordTo := hashBucketToAssignWordTo(word, 12)
		// tally a word for that bucket
		buckets[hashBucketToAssignWordTo]++

		// tally a word for that bucket in the map
		bucketsMap[hashBucketToAssignWordTo]++

		// increase tally for that hash=letter-bucket - num of words starting with that letter
		hashLetterBuckets[int(word[0])]++
	}

	// show how many words are in each bucket
	fmt.Println("as buckets:", buckets)
	fmt.Println("as bucketsMap:", bucketsMap)

	for i := 28; i <= 126; i++ {
		fmt.Printf("as hash letter buckets: %c, %v \n", i, hashLetterBuckets[i])
	}
}

func hashBucketToAssignWordTo(word string, numBuckets int) int {
	letter := int(word[0])
	// fmt.Printf("Letter is: %v", letter)
	// modulus to even out the distribution
	bucket := letter % numBuckets
	return bucket
}

func mobyDick() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	// http://www.gutenberg.org/files/2701/2701-0.txt
	if err != nil {
		log.Fatal("There was an error retrieving", err)
	}

	byteStream, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal("There was an error reading", err)
	}

	fmt.Printf("%s", byteStream)
}

func loremReading() {
	lorem := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
  eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
   veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea
    commodo consequat. Duis aute irure dolor in reprehenderit in voluptate
     velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat
      cupidatat non proident, sunt in culpa qui officia deserunt mollit anim
       id est laborum.`
	scanner := bufio.NewScanner(strings.NewReader(lorem))

	// count the words
	// fmt.Printf("scanner.Scan() returns a (word by word) %T \n", scanner.Scan())
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

}
