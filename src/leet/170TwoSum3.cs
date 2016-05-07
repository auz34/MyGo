public class TwoSum {
    
    private Dictionary<int, int> dict = new Dictionary<int, int>();

    // Add the number to an internal data structure.
	public void Add(int number) {
	    if (dict.ContainsKey(number)) {
	        dict[number]++;
	    } else {
	        dict[number] = 1;
	    }
	}

    // Find if there exists any pair of numbers which sum is equal to the value.
	public bool Find(int value) {
	    foreach (var kvp in dict) {
	        var target = value - kvp.Key;
	        if (target == kvp.Key) {
	            if (kvp.Value > 1) {
	                return true;
	            }
	        } else if (dict.ContainsKey(target)) {
	            return true;
	        }
	    }
	    
	    return false;
	}
}


// Your TwoSum object will be instantiated and called as such:
// TwoSum twoSum = new TwoSum();
// twoSum.Add(number);
// twoSum.Find(value);