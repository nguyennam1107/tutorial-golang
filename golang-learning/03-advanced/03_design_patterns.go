package main

import (
	"fmt"
	"sync"
)

// ============================================
// 1. SINGLETON PATTERN
// ============================================

type Database struct {
	connection string
}

var (
	dbInstance *Database
	dbOnce     sync.Once
)

func GetDatabaseInstance() *Database {
	dbOnce.Do(func() {
		fmt.Println("Creating database instance...")
		dbInstance = &Database{
			connection: "postgresql://localhost:5432",
		}
	})
	return dbInstance
}

// ============================================
// 2. FACTORY PATTERN
// ============================================

type Vehicle interface {
	Drive() string
}

type Car struct {
	Brand string
}

func (c Car) Drive() string {
	return fmt.Sprintf("Driving %s car", c.Brand)
}

type Bike struct {
	Type string
}

func (b Bike) Drive() string {
	return fmt.Sprintf("Riding %s bike", b.Type)
}

func CreateVehicle(vehicleType string) Vehicle {
	switch vehicleType {
	case "car":
		return Car{Brand: "Toyota"}
	case "bike":
		return Bike{Type: "Mountain"}
	default:
		return nil
	}
}

// ============================================
// 3. ABSTRACT FACTORY PATTERN
// ============================================

// GUI Factory
type Button interface {
	Render() string
}

type Checkbox interface {
	Check() string
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Windows implementation
type WindowsButton struct{}

func (wb WindowsButton) Render() string {
	return "Windows Button"
}

type WindowsCheckbox struct{}

func (wc WindowsCheckbox) Check() string {
	return "Windows Checkbox"
}

type WindowsFactory struct{}

func (wf WindowsFactory) CreateButton() Button {
	return WindowsButton{}
}

func (wf WindowsFactory) CreateCheckbox() Checkbox {
	return WindowsCheckbox{}
}

// Mac implementation
type MacButton struct{}

func (mb MacButton) Render() string {
	return "Mac Button"
}

type MacCheckbox struct{}

func (mc MacCheckbox) Check() string {
	return "Mac Checkbox"
}

type MacFactory struct{}

func (mf MacFactory) CreateButton() Button {
	return MacButton{}
}

func (mf MacFactory) CreateCheckbox() Checkbox {
	return MacCheckbox{}
}

// ============================================
// 4. BUILDER PATTERN
// ============================================

type House struct {
	Windows int
	Doors   int
	Floors  int
	HasGarage bool
	HasGarden bool
}

type HouseBuilder struct {
	house House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (hb *HouseBuilder) Windows(count int) *HouseBuilder {
	hb.house.Windows = count
	return hb
}

func (hb *HouseBuilder) Doors(count int) *HouseBuilder {
	hb.house.Doors = count
	return hb
}

func (hb *HouseBuilder) Floors(count int) *HouseBuilder {
	hb.house.Floors = count
	return hb
}

func (hb *HouseBuilder) WithGarage() *HouseBuilder {
	hb.house.HasGarage = true
	return hb
}

func (hb *HouseBuilder) WithGarden() *HouseBuilder {
	hb.house.HasGarden = true
	return hb
}

func (hb *HouseBuilder) Build() House {
	return hb.house
}

// ============================================
// 5. PROTOTYPE PATTERN
// ============================================

type Cloneable interface {
	Clone() Cloneable
}

type Document struct {
	Title   string
	Content string
	Author  string
}

func (d *Document) Clone() Cloneable {
	return &Document{
		Title:   d.Title,
		Content: d.Content,
		Author:  d.Author,
	}
}

// ============================================
// 6. ADAPTER PATTERN
// ============================================

// Old interface
type LegacyPrinter interface {
	PrintOldWay(text string)
}

type OldPrinter struct{}

func (op OldPrinter) PrintOldWay(text string) {
	fmt.Println("[OLD]", text)
}

// New interface
type ModernPrinter interface {
	Print(text string)
}

// Adapter
type PrinterAdapter struct {
	oldPrinter LegacyPrinter
}

func (pa PrinterAdapter) Print(text string) {
	pa.oldPrinter.PrintOldWay(text)
}

// ============================================
// 7. DECORATOR PATTERN
// ============================================

type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (sc SimpleCoffee) Cost() float64 {
	return 5.0
}

func (sc SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Milk decorator
type MilkDecorator struct {
	coffee Coffee
}

func (md MilkDecorator) Cost() float64 {
	return md.coffee.Cost() + 1.5
}

func (md MilkDecorator) Description() string {
	return md.coffee.Description() + ", Milk"
}

// Sugar decorator
type SugarDecorator struct {
	coffee Coffee
}

func (sd SugarDecorator) Cost() float64 {
	return sd.coffee.Cost() + 0.5
}

func (sd SugarDecorator) Description() string {
	return sd.coffee.Description() + ", Sugar"
}

// ============================================
// 8. FACADE PATTERN
// ============================================

type CPU struct{}

func (c CPU) Start() {
	fmt.Println("CPU started")
}

type Memory struct{}

func (m Memory) Load() {
	fmt.Println("Memory loaded")
}

type HardDrive struct{}

func (hd HardDrive) Read() {
	fmt.Println("HardDrive read")
}

// Facade
type ComputerFacade struct {
	cpu       CPU
	memory    Memory
	hardDrive HardDrive
}

func (cf ComputerFacade) Start() {
	fmt.Println("Starting computer...")
	cf.cpu.Start()
	cf.memory.Load()
	cf.hardDrive.Read()
	fmt.Println("Computer started!")
}

// ============================================
// 9. PROXY PATTERN
// ============================================

type Image interface {
	Display()
}

type RealImage struct {
	filename string
}

func (ri *RealImage) loadFromDisk() {
	fmt.Println("Loading image from disk:", ri.filename)
}

func (ri *RealImage) Display() {
	fmt.Println("Displaying:", ri.filename)
}

// Proxy
type ProxyImage struct {
	filename  string
	realImage *RealImage
}

func (pi *ProxyImage) Display() {
	if pi.realImage == nil {
		pi.realImage = &RealImage{filename: pi.filename}
		pi.realImage.loadFromDisk()
	}
	pi.realImage.Display()
}

// ============================================
// 10. OBSERVER PATTERN
// ============================================

type Observer interface {
	Update(message string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Notify(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

type ConcreteObserver struct {
	name string
}

func (co ConcreteObserver) Update(message string) {
	fmt.Printf("[%s] Received: %s\n", co.name, message)
}

// ============================================
// 11. STRATEGY PATTERN
// ============================================

type SortStrategy interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (bs BubbleSort) Sort(data []int) []int {
	fmt.Println("Sorting using Bubble Sort")
	// Implementation...
	return data
}

type QuickSort struct{}

func (qs QuickSort) Sort(data []int) []int {
	fmt.Println("Sorting using Quick Sort")
	// Implementation...
	return data
}

type Sorter struct {
	strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(data []int) []int {
	return s.strategy.Sort(data)
}

// ============================================
// 12. COMMAND PATTERN
// ============================================

type Command interface {
	Execute()
	Undo()
}

type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Light is ON")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Light is OFF")
}

type TurnOnCommand struct {
	light *Light
}

func (c TurnOnCommand) Execute() {
	c.light.TurnOn()
}

func (c TurnOnCommand) Undo() {
	c.light.TurnOff()
}

type TurnOffCommand struct {
	light *Light
}

func (c TurnOffCommand) Execute() {
	c.light.TurnOff()
}

func (c TurnOffCommand) Undo() {
	c.light.TurnOn()
}

type RemoteControl struct {
	command Command
}

func (rc *RemoteControl) SetCommand(command Command) {
	rc.command = command
}

func (rc *RemoteControl) PressButton() {
	rc.command.Execute()
}

func (rc *RemoteControl) PressUndo() {
	rc.command.Undo()
}

// ============================================
// EXAMPLES
// ============================================

func main() {
	fmt.Println("=== SINGLETON ===")
	db1 := GetDatabaseInstance()
	db2 := GetDatabaseInstance()
	fmt.Println("Same instance?", db1 == db2)

	fmt.Println("\n=== FACTORY ===")
	car := CreateVehicle("car")
	bike := CreateVehicle("bike")
	fmt.Println(car.Drive())
	fmt.Println(bike.Drive())

	fmt.Println("\n=== ABSTRACT FACTORY ===")
	var factory GUIFactory
	factory = WindowsFactory{}
	btn := factory.CreateButton()
	cb := factory.CreateCheckbox()
	fmt.Println(btn.Render())
	fmt.Println(cb.Check())

	fmt.Println("\n=== BUILDER ===")
	house := NewHouseBuilder().
		Windows(6).
		Doors(2).
		Floors(2).
		WithGarage().
		WithGarden().
		Build()
	fmt.Printf("House: %+v\n", house)

	fmt.Println("\n=== DECORATOR ===")
	coffee := SimpleCoffee{}
	fmt.Printf("%s: $%.2f\n", coffee.Description(), coffee.Cost())

	coffeeWithMilk := MilkDecorator{coffee: coffee}
	fmt.Printf("%s: $%.2f\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())

	coffeeWithMilkAndSugar := SugarDecorator{coffee: coffeeWithMilk}
	fmt.Printf("%s: $%.2f\n", coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())

	fmt.Println("\n=== FACADE ===")
	computer := ComputerFacade{}
	computer.Start()

	fmt.Println("\n=== PROXY ===")
	image := &ProxyImage{filename: "photo.jpg"}
	image.Display() // Loading + Display
	image.Display() // Display only (cached)

	fmt.Println("\n=== OBSERVER ===")
	subject := &Subject{}
	subject.Attach(ConcreteObserver{name: "Observer1"})
	subject.Attach(ConcreteObserver{name: "Observer2"})
	subject.Notify("Hello Observers!")

	fmt.Println("\n=== STRATEGY ===")
	sorter := &Sorter{}
	data := []int{5, 2, 8, 1, 9}

	sorter.SetStrategy(BubbleSort{})
	sorter.Sort(data)

	sorter.SetStrategy(QuickSort{})
	sorter.Sort(data)

	fmt.Println("\n=== COMMAND ===")
	light := &Light{}
	remote := &RemoteControl{}

	onCommand := TurnOnCommand{light: light}
	remote.SetCommand(onCommand)
	remote.PressButton()
	remote.PressUndo()

	offCommand := TurnOffCommand{light: light}
	remote.SetCommand(offCommand)
	remote.PressButton()
	remote.PressUndo()
}
