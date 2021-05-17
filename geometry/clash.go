package geometry

func sphereXsphere(s1, s2 *Sphere) (clash bool) {

	d := Sub(s1.centre, s2.centre).Len()
	sum := (s1.radius + s2.radius)

	if distEq(sum, d) {
		// TODO handle touching
		return false
	}

	clash = d <= (s1.radius + s2.radius)
	return

}

func sphereXgeom(s1 *Sphere, g2 Geometry) (clash bool, err error) {
	switch g := g2.(type) {
	case *Sphere:
		clash = sphereXsphere(s1, g)
		return
	}

	err = ErrNotImplemented
	return

}

func clash(g1, g2 Geometry) (clash bool, err error) {

	switch g := g1.(type) {
	case *Sphere:
		sphereXgeom(g, g2)
	}

	err = ErrNotImplemented
	return
}
