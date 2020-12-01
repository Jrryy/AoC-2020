<?php
# Part 1:
$data = array_map('intval', file("input.txt", 2));
rsort($data);
foreach($data as $value) {
	$i = count($data) - 1;
	while($value + $data[$i] < 2020) $i -= 1;
	if ($value + $data[$i] == 2020) {
		echo $value*$data[$i];
		break;
	}
}

echo "\n---------------------\n";

# Part 2:
foreach($data as $value) {
	$i = count($data) - 1;
	while($value + $data[$i] < 2020) {
		$j = $i - 1;
		while($value + $data[$i] + $data[$j] < 2020) $j -= 1;
		if ($value + $data[$i] + $data[$j] == 2020) {
			echo $value*$data[$i]*$data[$j];
			exit();
		}
		$i -= 1;
	}
}
?>
