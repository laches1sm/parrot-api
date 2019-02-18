package design

import . "goa.design/goa/dsl"

var _ = API("parrot", func() {
	Title("service-parrot")
	Description("it's parrots")
})

var _ = Service("parrot", func() {
	Description("parrots")
	Method("add-parrot", func() {
		Payload(func() {
			Attribute("name", String, "Parrot name")
			Attribute("colour", String, "Parrot colour")
			Attribute("breed", String, "Parrot breed")
			Required("name", "colour", "breed")
		})
		Result(Int)
		HTTP(func() {
			POST("post-parrot/{name}/{breed}/{colour}")
			Response(StatusCreated)
		})
	})
})
