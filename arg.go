package ncrack

import "strconv"

type arg struct {
	Val interface{}
}

func (c arg) Arg() string {
	switch ac := c.Val.(type) {
	case uint16:
		return strconv.Itoa(int(ac))
	case string:
		return ac
	case ServiceOptions:
		return ac.String()
	}
	return ""
}

func (c arg) withModificator(fn func(string) string) string {
	if val := c.Arg(); val != "" {
		return fn(val)
	}
	return ""
}

func (c arg) withModificatorPassAnyway(fn func(string) string) string {
	return fn(c.Arg())
}

func (c arg) WithCommaAfter() string {
	return c.withModificator(func(a string) string {
		return a + ","
	})
}

func (c arg) WithCommaBefore() string {
	return c.withModificator(func(a string) string {
		return "," + a
	})
}

func (c arg) WithSchemaAfter() string {
	return c.withModificator(func(a string) string {
		return a + "://"
	})
}

func (c arg) WithColonAfter() string {
	return c.withModificator(func(a string) string {
		return a + ":"
	})
}

func (c arg) WithColonBetween(a2 string) string {
	return c.withModificatorPassAnyway(func(a1 string) string {
		return withSymbolBetween(a1, a2, ":")
	})
}

func (c arg) WithColumnBetween(a2 string) string {
	return c.withModificatorPassAnyway(func(a1 string) string {
		return withSymbolBetween(a1, a2, ",")
	})
}

func withSymbolBetween(a1, a2, sym string) string {
	if a1 != "" && a2 != "" {
		return a1 + sym + a2
	}
	if a1 == "" && a2 != "" {
		return a2
	}
	if a1 != "" && a2 == "" {
		return a1
	}
	return ""
}
