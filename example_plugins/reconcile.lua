function process(input_str, additional_params)
	local pattern = (additional_params and additional_params[1]) or "(?i).+"

	local counts = {}
	local order = {}
	for match in re.gmatch(input_str, pattern) do
		if not counts[match] then
			counts[match] = 0
			table.insert(order, match)
		end

		counts[match] = counts[match] + 1
	end

	local output_lines = {}
	for _, match in ipairs(order) do
		local count = counts[match]
		if count > 1 then
			table.insert(output_lines, count .. ": " .. match)
		end
	end

	return table.concat(output_lines, "\n")
end
