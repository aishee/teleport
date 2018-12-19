import React from 'react';
import { storiesOf } from '@storybook/react'
import { withInfo } from '@storybook/addon-info'
import Box from './Box';

storiesOf('Box', module)
  .addDecorator(withInfo)
  .addParameters({
    info: {
      inline: true,
      header: false,
    },
  })
  .add('Layout component', () => (
    <Box color="white" bg="blue">
      this is a simple div
    </Box>
  ))
  .add('Padding', () => (
    <Box bg="blue" p={3}>
      Hello
    </Box>
  ))
  .add('Margin', () => (
    <Box bg="blue" m={3}>
      Hello
    </Box>
  ))
  .add('Color', () => (
    <Box bg="blue" p={3} color='white'>
      Hello
    </Box>
  ))
  .add('Background Color', () => (
    <Box bg="blue" p={3} color='white'>
      Hello
    </Box>
  ))
  .add('Width', () => (
    <Box
      p={3}
      width={1 / 2}
      color='white'
      bg='blue'>
      Half Width
    </Box>
  ))
  .add('Pixel Width', () => (
    <Box
      p={3}
      width={256}
      color='white'
      bg='blue'>
      256px width
    </Box>
  ))
  .add('VW Width', () => (
    <Box
      p={3}
      width='50vw'
      color='white'
      bg='blue'>
      50vw width
    </Box>
  ))
  .add('Directional Padding', () => (
    <Box p={3}>
      <Box m={1} pt={3} color='white' bg='blue'>Padding Top</Box>
      <Box m={1} pr={3} color='white' bg='blue'>Padding Right</Box>
      <Box m={1} pb={3} color='white' bg='blue'>Padding Bottom</Box>
      <Box m={1} pl={3} color='white' bg='blue'>Padding Left</Box>
      <Box m={1} px={3} color='white' bg='blue'>Padding X-Axis</Box>
      <Box m={1} py={3} color='white' bg='blue'>Padding Y-Axis</Box>
    </Box>
  ))
  .add('Directional Margin', () => (
    <Box p={3}>
      <Box mt={3} color='white' bg='blue'>Margin Top</Box>
      <Box mr={3} color='white' bg='blue'>Margin Right</Box>
      <Box mb={3} color='white' bg='blue'>Margin Bottom</Box>
      <Box ml={3} color='white' bg='blue'>Margin Left</Box>
      <Box mx={3} color='white' bg='blue'>Margin X-Axis</Box>
      <Box my={3} color='white' bg='blue'>Margin Y-Axis</Box>
    </Box>
  ));