package main

import "fmt"


type Point struct {
  X int
  Y int
}

func (p *Point) Move(dx int,dy int){
  p.X+=dx
  p.Y+=dy
}

type Square struct {
  center Point
  length int 
}

func (s *Square)Move(dx int,dy int){
  s.center.Move(dx, dy)
}

func (s Square)Area() int{
  return s.length*s.length
}

func NewSquare(x int, y int, length int) (*Square, error){
  if length < 0 {
    return nil, fmt.Errorf("Negative length is invalid.")
  }
  center := Point{x,y}
  square := Square{center,length}

  return &square,nil
}

func (s Square) String() string {
        return fmt.Sprintf("Center is in %v ,%v, Area is %v\n", s.center.X, s.center.Y, s.Area())
}


func main() {
	square_a, err := NewSquare(40,-50,12)
  if err!=nil{
    fmt.Printf("Error: %v\n",err)
  }
  
  fmt.Printf(square_a.String())
  square_a.Move(4, -7)
  fmt.Printf(square_a.String())

  _ , err = NewSquare(40,-50,-12)
  if err!=nil{
    fmt.Printf("Error: %v\n",err)
  }
  
  
}
