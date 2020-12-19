using System;
using Xunit;

namespace LeetCodeProblems
{
    public class TwoSum
    {
        public int[] TwoSumSimple(int[] nums, int target)
        {
            for (var i = 0; i < nums.Length; i++)
            {
                for (var j = i + 1; j < nums.Length; j++)
                {
                    if (nums[i] + nums[j] == target)
                        return new[] {i, j};
                }
            }

            throw new ArgumentException();
        }

        [Theory]
        [InlineData(new[] {3, 2, 4}, 6, new[] {1, 2})]
        [InlineData(new[] {2, 7, 11, 15}, 9, new[] {0, 1})]
        [InlineData(new[] {3, 3}, 6, new[] {0, 1})]
        public void TwoSumTest(int[] nums, int target, int[] expected)
        {
            var result = TwoSumSimple(nums, target);
            Assert.Equal(expected, result);
        }
    }
}