module.exports = {
  env: {
    browser: true,
    jquery: true,
    es2021: true,
  },
  extends: ['plugin:react/recommended', 'standard'],
  globals: {
    RequestInit: true,
    JSX: true,
  },
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 'latest',
    sourceType: 'module',
  },
  plugins: ['react', '@typescript-eslint'],
  rules: {
    'no-tabs': 0,
    'no-mixed-spaces-and-tabs': 0,
    'react/jsx-uses-react': 'off',
    'react/react-in-jsx-scope': 'off',
    'no-unused-vars': ['error', { varsIgnorePattern: '^_$', }],
    'comma-dangle': [
      'error',
      {
        arrays: 'never',
        objects: 'ignore',
        imports: 'never',
        exports: 'never',
        functions: 'never',
      }
    ],
  },
}
