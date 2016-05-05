/*Given an array of words and a length L, format the text such that each line has exactly L characters and is fully
(left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces
' ' when necessary so that each line has exactly L characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide
evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left justified and no extra space is inserted between words.

For example,
words: ["This", "is", "an", "example", "of", "text", "justification."]
L: 16.

Return the formatted lines as:
[
"This    is    an",
"example  of text",
"justification.  "
]
Note: Each word is guaranteed not to exceed L in length.*/

/*C# Solution

Use cases to keep in mind
//var words = new List<string> { "This", "is", "an", "example", "of", "text", "justification." };
            //var maxWidth = 16;

            //var words = new List<string> { "" };
            //var maxWidth = 0;

            //var words = new List<string> { "a", "b", "c", "d", "e" };
            //var maxWidth = 1;


public class Solution
    {
        private List<List<string>> PrepareLines(string[] words, int maxWidth)
        {
            var preResult = new List<List<string>>();
            var curLine = new List<string>();
            var curLength = 0;
            foreach (var word in words)
            {
                var futureLength = curLength + (curLength > 0 ? 1 : 0) + word.Length;
                if (futureLength <= maxWidth)
                {
                    curLine.Add(word);
                    curLength = futureLength;
                }
                else
                {
                    preResult.Add(curLine);
                    curLine = new List<string>();
                    curLine.Add(word);
                    curLength = word.Length;
                }
            }

            if (curLength > 0)
            {
                preResult.Add(curLine);
            }

            return preResult;
        }

        private string JustifyLine(List<string> line, int maxWidth)
        {
            var alphaLength = line.Sum(word => word.Length);
            var betweenSpaces = line.Count - 1;
            var extraSpaces = maxWidth - alphaLength - betweenSpaces;
            var extraPerEachLine = 0;
            var extraRemainder = 0;
            if (extraSpaces > 0 && betweenSpaces > 0)
            {
                extraPerEachLine = extraSpaces / betweenSpaces;
                extraRemainder = extraSpaces % betweenSpaces;
            }


            var sb = new StringBuilder();
            for (var i = 0; i < line.Count - 1; i++)
            {
                sb.Append(line[i]);
                sb.Append("".PadRight(extraPerEachLine));

                if (extraRemainder > 0)
                {
                    sb.Append(" ");
                    extraRemainder--;
                }
                sb.Append(" ");
            }

            sb.Append(line[line.Count - 1]);
            if (sb.Length < maxWidth)
            {
                sb.Append("".PadRight(maxWidth - sb.Length));
            }
            return sb.ToString();
        }

        public IList<string> FullJustify(string[] words, int maxWidth)
        {
            var result = new List<string>();
            var preResult = this.PrepareLines(words, maxWidth);
            if (!preResult.Any())
            {
                result.Add("".PadRight(maxWidth));
                return result;
            }

            for (var i = 0; i < preResult.Count - 1; i++)
            {
                result.Add(this.JustifyLine(preResult[i], maxWidth));
            }

            var lastLine = string.Join(" ", preResult[preResult.Count - 1]);
            if (lastLine.Length < maxWidth)
            {
                lastLine += "".PadRight(maxWidth - lastLine.Length);
            }
            result.Add(lastLine);

            return result;
        }
    }
 */

