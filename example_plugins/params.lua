function process(input, parms)
	ps = type(parms) == "table" and table.concat(parms, ", ") or "none"
	return "input:" .. input .. "\nparams:" .. ps
end

