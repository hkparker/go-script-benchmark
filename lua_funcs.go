package script

var lua_modulo_test = `
	for i=1,1000000 do
		if (math.fmod(i,30)<1) then
			local x = 1
		end
	end
`

var lua_recursion_test = `
	function foo (i)
		if i < 100 then
			foo(i+1)
		else
			return
		end
	end
	foo(0)
`
