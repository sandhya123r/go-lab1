
package main

import "fmt"

type LRUCache struct{

  Hashtable map[int]*node
  size int
  count int
  ll *list

}


type node struct{
  key int
  value int
  next *node
  prev *node
}

type list struct{
  head *node
  tail *node
}

func addnode(l *list,n *node){
  if(l==nil){
    newlist:=new(list)
    l:=newlist
    l.head=n
    l.head.prev=nil
    l.tail=l.head
    l.head.next=nil
    return
  }
  n.prev=l.head
  n.next=l.head.next
  l.head.next.prev=n
  l.head.next=n
}

func removenode(l *list,n *node){
  var prev_node *node=nil
  var next_node *node=nil
  prev_node=n.prev
  next_node=n.next
  
  prev_node.next=next_node
  next_node.prev=prev_node
  
}

func movetohead(l *list, n *node){
  if(l==nil){
    return
  }
  removenode(l,n)
  addnode(l,n)
}

func pop(l *list)*node{ 
  popped_node:=l.tail
  fmt.Println(popped_node.key)
  if(l.tail!=nil && l.tail.prev!=nil){
    l.tail=l.tail.prev
  }
  removenode(l,popped_node)
  return popped_node
}

func Cache(cache *LRUCache,size int){

  cache.count=0
  cache.size=size
  
  var head *node=nil
  head.prev=nil
  
  var tail *node=nil
  tail.next=nil
  
  head.prev=tail
  tail.next=head
  
}

func get(cache *LRUCache,key int)int{
  
  hash:=cache.Hashtable
  targetnode,flag:=hash[key]
  if(!flag){
    return -1
  }
  
  linkedlist:=cache.ll
  movetohead(linkedlist,targetnode)
  
  return targetnode.value
  
}

func set(cache *LRUCache,key int, value int){
  
  hash:=cache.Hashtable
  targetnode,flag:=hash[key]
  if(!flag){
    newnode:=new (node)
    newnode.key=key
    newnode.value=value
    if(len(cache.Hashtable)==0){
      Hashtable:=make(map[int]*node)
      cache.Hashtable=Hashtable
    }
    cache.Hashtable[key]=newnode
    addnode(cache.ll,newnode)
    cache.count++
    if(cache.count>cache.size){
      fmt.Println("Count > size ")
      tailnode:=pop(cache.ll)
      _,ok:=hash[tailnode.key]
      if(!ok){
        delete(cache.Hashtable,tailnode.key)
      }
      cache.count--
    }
  } else {
    targetnode.value=value
    movetohead(cache.ll,targetnode)
  }
}

func main(){
  
  l:=new(LRUCache)
  l.size=100
  fmt.Println("Default LRU Cache",l)
  set(l,2,2)
  set(l,3,5)
  set(l,3,9)
  set(l,4,5)
  set(l,10,5)
  set(l,300,5)
  fmt.Println(get(l,2))
}

