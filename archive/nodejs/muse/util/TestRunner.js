let itemIndex = 0;

export const pick = (func) => {
  if (func()) {
    console.log('✓ Item [%d] PASSED', itemIndex);
  } else {
    console.error('✗ Item [%d] FAILED', itemIndex);
  }

  itemIndex++;
};
