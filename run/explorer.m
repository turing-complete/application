function explorer()
  values = h5read('explorer.h5', '/values');
  assert(size(values, 2) == 1);
  histogram(values, 100, 'Normalization', 'pdf');
  title('Empirical PDF');
  xlabel('Uncertain parameter');
  ylabel('Probability density');
end
