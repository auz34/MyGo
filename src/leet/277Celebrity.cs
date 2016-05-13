using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Celebrity
{
    class Program
    {
        static void Main(string[] args)
        {
        }
    }

    /* The Knows API is defined in the parent class Relation.
      bool Knows(int a, int b); */

    public class Relation
    {
        protected bool Knows(int a, int b)
        {
            return true;
        }
    }

    public class Solution : Relation
    {
        public int FindCelebrity(int n)
        {
            var candidates = new HashSet<int>();
            for (var c = 0; c < n; c++)
            {
                var knowsSomeone = false;
                // we need to make sure this candidate don't know anyone
                for (var i = 0; i < n; i++)
                {
                    if (c == i)
                    {
                        continue;
                    }

                    knowsSomeone |= this.Knows(c, i);

                    if (knowsSomeone)
                    {
                        break;
                    }
                }

                if (!knowsSomeone)
                {
                    candidates.Add(c);
                }
            }

            foreach (var candidate in candidates)
            {
                var isKnown = true;
                // we need to make sure everybody else know the candidate
                for (var i = 0; i < n; i++)
                {
                    if (candidate == i)
                    {
                        continue;
                    }

                    isKnown &= this.Knows(i, candidate);

                    if (!isKnown)
                    {
                        break;
                    }
                }

                if (isKnown)
                {
                    return candidate;
                }
            }

            return -1;
        }
    }
}
