package script

var javascript_modulo_test = `
	for (i = 0; i<1000000; i++) {
		if (i%30 < 1) {
			var x = 1
		}
	}
`

var javascript_recursion_test = `
		foo = function(i) {
			if (i < 100) {
				foo(i+1)
			} else {
				return
			}
				
		}
		foo(0)
`
