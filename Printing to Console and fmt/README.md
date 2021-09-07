# FMT Cheatsheet in Go

## General

- %v (value in default format)
- %T (type)
- %% (literal %) - to print a percent sign

## Boolean

- %t (true or false)

## Integer

- %b (base 2)
- %o (base 8)
- %d (base 10)
- %x (base 16) or %X (base 16 but result in upper case)

## Floating Points

- %e (scientific notation)
- %f / %F (decimal no exponent) - cuts off some decimal places
- %g (for large exponents) - shows all decimal points

### Strings

- %s (default)
- %q (double quoted string)

## Width and Precision

- %f (default width, default precision)
- %9f (width 9, default precision)
- %.2f (default width, precision 2)
- %9.2f (width 9, precision 2)
- %9.f (width 9, precision 0)

## Padding

- %09d (pads digit with preceeding zeros to make length 9)
- %-9d (Pads with spaces i.e. width 9, left-justified)

## Methods

- Sprintf() - format without printing
- Printf() - format with printing
