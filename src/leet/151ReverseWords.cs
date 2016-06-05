using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ReverseWords
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            Console.WriteLine(s.ReverseWords("   a   b "));
            Console.WriteLine(s.ReverseWords("the sky is blue"));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        public string ReverseWords(string s)
        {
            var words = s.Split(' ')
                .Where(si => !string.IsNullOrEmpty(si))
                .ToList();
            
            var mid = words.Count / 2;
            for (var i = 0; i < mid; i++)
            {
                var inter = words[i];
                words[i] = words[words.Count - i - 1];
                words[words.Count - i - 1] = inter;
            }

            return string.Join(" ", words);
        }
    }
}
