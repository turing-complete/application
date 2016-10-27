function depict(file)
  values = h5read(file, '/values');
  assert(size(values, 2) == 1);
  histogram(values, 100, 'Normalization', 'pdf');
  title('Empirical PDF');
  xlabel('Total energy (J)');
  ylabel('Probability density');
end
