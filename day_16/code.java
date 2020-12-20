import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.ArrayList;

class Code {
	public static void main(String[] args) throws FileNotFoundException{
		ArrayList< int[] > ranges = new ArrayList< int[] >();
		int sum = 0;
		long departures = 1;
		Pattern pattern = Pattern.compile("^[\\w\\s]+: (\\d+)-(\\d+) or (\\d+)-(\\d+)$");
		File file = new File("input.txt");
		Scanner scanner = new Scanner(file);
		
		while (scanner.hasNextLine()){
			String line = scanner.nextLine();
			if (line.equals("")) break;
			Matcher matcher = pattern.matcher(line);
			matcher.find();
			ranges.add(new int[]{
				Integer.parseInt(matcher.group(1)),
				Integer.parseInt(matcher.group(2)),
				Integer.parseInt(matcher.group(3)),
				Integer.parseInt(matcher.group(4))
			});
		}

		int[] fields = new int[ranges.size()];
		for (int i = 0; i < fields.length; i++) {
			fields[i] = -1;
		}
		ArrayList<ArrayList<Integer>> possibilities = new ArrayList<ArrayList<Integer>>();
		for (int i = 0; i < fields.length; i++) {
			possibilities.add(new ArrayList<Integer>());
		}


		scanner.nextLine();
		String[] myTicket = scanner.nextLine().split(",");
		for (int i = 0; i < ranges.size(); i++) {
			int[] range = ranges.get(i);
			for (int j = 0; j < myTicket.length; j++) {
				int n = Integer.parseInt(myTicket[j]);
				if (n >= range[0] && n <= range[1] || n >= range[2] && n <= range[3]){
					possibilities.get(i).add(j);
				}
			}
		}

		scanner.nextLine();
		scanner.nextLine();
		
		while (scanner.hasNextLine()){
			String nearbyTicket = scanner.nextLine();
			String[] numbers = nearbyTicket.split(",");
			boolean valid = true;

			for (String number : numbers) {
				boolean inRange = false;
				int n = Integer.parseInt(number);
				for (int[] range : ranges) {
					if (n >= range[0] && n <= range[1] || n >= range[2] && n <= range[3]) {
						inRange = true;
						continue;
					}
				}
				if (!inRange) {
					valid = false;
					sum += n;
				}
			}
			
			if (valid) {
				for (int i = 0; i < ranges.size(); i++) {
					int[] range = ranges.get(i);
					for (int j = 0; j < numbers.length; j++) {
						int n = Integer.parseInt(numbers[j]);
						if (!(n >= range[0] && n <= range[1] || n >= range[2] && n <= range[3])){
							ArrayList<Integer> posRange = possibilities.get(i);
							if (posRange.contains(j)) posRange.remove(Integer.valueOf(j));
						}
					}
				}
			}
		}

		scanner.close();

		for (int i = 0; i < possibilities.size(); i++) {
			ArrayList<Integer> pos = possibilities.get(i);
			if (pos.size() == 1) {
				int value = pos.get(0);
				if (0 <= i && i <= 5) {
					departures *= Long.parseLong(myTicket[value]);
				}
				for (int j = 0; j < possibilities.size(); j++) {
					ArrayList<Integer> posRange = possibilities.get(j);
					if (posRange.contains(value)) posRange.remove(Integer.valueOf(value));
				}
				i = -1;
			}
		}

		System.out.println(sum);
		System.out.println(departures);
	}
}
