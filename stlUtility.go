import (
	"github.com/hschendel/stl"
)

type sortableTriangle struct {
  Triangle *stl.Triangle // pointer to the tri
  Group int // this will track which group the triangle is part of
}

// set up sortableTriangles
func toSortable (tris []stl.Triangle) []sortableTriangle {
  var ret [len(tris)]sortableTriangle // allocate array to return

  for i := 0; i < len(tris); i++ {
    ret[i].Group = i // set the group to the current numbe, each tri needs a different number
    ret[i].Triangle = &tris[i]
  }
  return ret
}

// test 2 triangles for connectedNess. 2 are connected if they share an edge (2 points)
func doTrianglesMeet (triangle1 *stl.Triangle, triangle2 *stl.Triangle) bool {
  count := 0
  tri1 := *triangle1
  tri2 := *triangle2
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if tri1.Vertices[i][0] == tri2.Vertices[j][0] && tri1.Vertices[i][1] == tri2.Vertices[j][1] && tri1.Vertices[i][2] == tri2.Vertices[j][2] {
        count++
        if count > 1 {
          return true // return true if the count goes above 1 shared point
        }
      }
    }
  }
  return false // return false if we have not found enough shared points
}

// iterate over the sortable triangles, set the group number accordingly
func groupSortables (sTris *[]sortableTriangle) int{
  changed := true // we want to stop when we detect no more changes
  for changed == true {
    changed = false // we should stop if we make no changes this cycle
    for i := 0; i < len(*sTris); i++ {
      for j := i + 1; j < len(*sTris); j++ {
        if doTrianglesMeet((*sTris)[i].Triangle, (*sTris)[j].Triangle) {
          // these triangles share an edge
          if (*sTris)[i].Group < (*sTris)[j].Group {
            // the groups are not set, i should be less than j if they share an edge
            changed = true // we will need to iterate again
            (*sTris)[j].Group = (*sTris)[i].Group
          }
        }
      }
    }
  }
  return (*sTris)[len(sTris) - 1].Group // return the number of groups found. this is equal to group number of the last triangle  
}
