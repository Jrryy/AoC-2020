counter_1 = 0
counter_2 = 0
mandatory_fields = {"byr" => "", "iyr" => "", "eyr" => "", "hgt" => "", "hcl" => "", "ecl" => "", "pid" => ""}

passport_fields = mandatory_fields.clone

def correct_value?(field, value)
	case field
	when "byr"
		return /^[0-9]{4}$/.match?(value) && value.to_i.between?(1920, 2002)
	when "iyr"
		return /^[0-9]{4}$/.match?(value) && value.to_i.between?(2010, 2020)
	when "eyr"
		return /^[0-9]{4}$/.match?(value) && value.to_i.between?(2020, 2030)
	when "hgt"
		/^[0-9]{2,3}(cm|in)$/.match(value) do |m|
			case m[1]
			when "cm" then return value.to_i.between?(150, 193)
			when "in" then return value.to_i.between?(59, 76)
			end
		end
	when "hcl"
		return /^#[0-9a-f]{6}$/.match?(value)
	when "ecl"
		return /^(amb|blu|brn|gry|grn|hzl|oth)$/.match?(value)
	when "pid"
		return /^[0-9]{9}$/.match?(value)
	else
		return true
	end
end

File.foreach("input.txt") {
	|line|
	if line == "\n"
		if passport_fields.all? { |key, value| value != "" }
			counter_1 += 1
			if passport_fields.all? { |key, value| correct_value?(key, value) == true }
				counter_2 += 1
			end
		end
		passport_fields = mandatory_fields.clone
	else
		line.split(" ").each {
			|field|
			key, value = field.split(":")
			passport_fields[key] = value
		}
	end
}

if passport_fields.all? { |key, value| value != "" }
	counter_1 += 1
	if passport_fields.all? { |key, value| correct_value?(key, value) == true }
		counter_2 += 1
	end
end

puts counter_1
puts counter_2
