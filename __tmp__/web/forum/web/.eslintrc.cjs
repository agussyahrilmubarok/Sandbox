module.exports = {
  env: {
    browser: true,
    es2021: true,
    'cypress/globals': true
  },
  settings: {
    react: {
      version: 'detect'
    }
  },
  extends: ['standard', 'plugin:cypress/recommended', 'plugin:react/recommended', 'plugin:storybook/recommended'],
  overrides: [
    {
      env: {
        node: true
      },
      files: ['.eslintrc.{js,cjs}'],
      parserOptions: {
        sourceType: 'script'
      }
    }
  ],
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module'
  },
  plugins: ['react'],
  rules: {
    'space-before-function-paren': [
      'error',
      {
        named: 'never',
        anonymous: 'always',
        asyncArrow: 'always'
      }
    ],
    'multiline-ternary': 'off'
  },
  ignorePatterns: ['dist/', 'node_modules/']
}
