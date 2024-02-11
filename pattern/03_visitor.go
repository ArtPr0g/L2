package pattern

/*
	 Этот паттерн применяется для разделения алгоритма от структуры объектов, так чтобы можно было добавить новые операции над объектами без изменения самих объектов. 
*/
import "fmt"

type Inspector interface {
    VisitOffice(office *Office)
    VisitHR(hr *HRDepartment)
    VisitIT(it *ITDepartment)
}

type TaxInspector struct {
    Name string
}

func (inspector *TaxInspector) VisitOffice(office *Office) {
    fmt.Printf("%s is inspecting the office\n", inspector.Name)
}

func (inspector *TaxInspector) VisitHR(hr *HRDepartment) {
    fmt.Printf("%s is inspecting the HR Department\n", inspector.Name)
}

func (inspector *TaxInspector) VisitIT(it *ITDepartment) {
    fmt.Printf("%s is inspecting the IT Department\n", inspector.Name)
}


type OfficeElement interface {
    Accept(inspector Inspector)
}

type Office struct {
    Sections []OfficeElement
}

func (o *Office) Accept(inspector Inspector) {
    inspector.VisitOffice(o)
    for _, section := range o.Sections {
        section.Accept(inspector)
    }
}

type HRDepartment struct{}

func (hr *HRDepartment) Accept(inspector Inspector) {
    inspector.VisitHR(hr)
}

type ITDepartment struct{}

func (it *ITDepartment) Accept(inspector Inspector) {
    inspector.VisitIT(it)
}

func main() {
    taxInspector := &TaxInspector{Name: "John Doe"}

    // Создайте офис с отделами кадров и ИТ
    office := &Office{
        Sections: []OfficeElement{&HRDepartment{}, &ITDepartment{}},
    }

    // Налоговый инспектор посещает офис и его отделы
    office.Accept(taxInspector)
}
