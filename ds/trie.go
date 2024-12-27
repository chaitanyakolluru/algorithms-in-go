package ds

// if its not a priority queue, its a trie They are pronounced like tree (its
// named after re"trie"val tree). So its a trie tree

// the easiest way to visualize a trie is to think of auto-complete.

// caching and auto-complete implementations mainly

// **

// c           m
//
// a r d       a   t  e
//
// t t l e     r
//
// s           s

// insertion(str)
// curr = head
// for c in str:
//   if curr[c]
// 	  curr = curr[c]
// 		else:
// 		  node =cN()
// 			curr[c] = node
// 			curr = node
//
// curr.isWord = true
