using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace RepeatedDNA
{
    class Program
    {
        static void Main(string[] args)
        {
        }
    }

    public class Solution
    {
        public IList<string> FindRepeatedDnaSequences(string s)
        {
            var dict = new Dictionary<string, int>();

            for (var i = 0; i < s.Length - 10 + 1; i++)
            {
                var seq = s.Substring(i, 10);
                if (dict.ContainsKey(seq))
                {
                    dict[seq]++;
                }
                else
                {
                    dict.Add(seq, 1);
                }
            }

            return new List<string>(dict.Where(kvp => kvp.Value > 1).Select(kvp => kvp.Key));
        }
    }
}
