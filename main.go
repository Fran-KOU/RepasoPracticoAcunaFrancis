package main

import "fmt"

const tarifaIVA = 0.15

type Producto struct {
	Codigo      string
	Descripcion string
	Precio      float64
	GravaIVA    bool //Si incluye o no IVA
}

type Catalogo struct {
	Productos map[string]Producto
}

type Linea struct {
	Producto Producto
	Cantidad int
}

type Factura struct {
	Numero  string
	Cliente string
	Lineas  []Linea
}

func (c Catalogo) Buscar(codigo string) (Producto, bool) {
	p, ok := c.Productos[codigo]
	return p, ok
}

func (f *Factura) AgregarLinea(p Producto, cantidad int) {
	f.Lineas = append(f.Lineas, Linea{Producto: p, Cantidad: cantidad})
}

func (l Linea) Subtotal() float64 {
	return l.Producto.Precio * float64(l.Cantidad)

}

func (f Factura) Subtotal() float64 {
	total := 0.0
	for _, linea := range f.Lineas {
		total += linea.Subtotal()
	}
	return total
}

func (f Factura) IVA() float64 {
	gravado := 0.0
	for _, l := range f.Lineas {
		if l.Producto.GravaIVA {
			gravado += l.Subtotal()
		}
	}
	return gravado * tarifaIVA
}

func (f Factura) Total() float64 {
	return f.Subtotal() + f.IVA()
}

func main() {
	fmt.Println("Minimarket La Floresta")
	fmt.Println("")
	cat := Catalogo{Productos: map[string]Producto{
		"P01": {Codigo: "P01", Descripcion: "Leche 1L", Precio: 1.00, GravaIVA: false},
		"P02": {Codigo: "P02", Descripcion: "Pan", Precio: 0.20, GravaIVA: false},
		"P03": {Codigo: "P03", Descripcion: "Gaseosa 3L", Precio: 2.80, GravaIVA: true},
		"P04": {Codigo: "P04", Descripcion: "Detergente 1kg", Precio: 3.20, GravaIVA: true},
	}}
	f := Factura{Numero: "001-001-000000124", Cliente: "Daniela Cevallos"}
	codigos := []string{"P01", "P02", "P03", "P04"}
	cantidades := []int{2, 10, 1, 1, 1}
	for i, codigo := range codigos {
		p, ok := cat.Buscar(codigo)
		if !ok {
			fmt.Printf("Aviso: el codigo %s no existe, por ende, no se factura\n", codigo)
			continue
		}
		f.AgregarLinea(p, cantidades[i])
	}
	fmt.Printf("Factura %s. Cliente: %s\n", f.Numero, f.Cliente)
	fmt.Printf("Cant    Descripción      P.Unitario   Subtotal \n")
	for _, l := range f.Lineas {
		fmt.Printf("%4d %-17s %8.2f %9.2f\n", l.Cantidad, l.Producto.Descripcion, l.Producto.Precio, l.Subtotal())
	}
	fmt.Printf("%-32s %9.2f\n", "Subtotal", f.Subtotal())
	fmt.Printf("%-32s %9.2f\n", "IVA", f.IVA())
	fmt.Printf("%-32s %9.2f\n", "Total", f.Total())

	p := cat.Productos["P03"]
	p.Precio = 3.00
	cat.Productos["P03"] = p
	fmt.Printf("La gaseosa sube a $%.2f en el catalogo. \n", cat.Productos["P03"].Precio)
	fmt.Printf("La factura %s la conserva a $%2.f. \n", f.Numero, f.Lineas[2].Producto.Precio)
}