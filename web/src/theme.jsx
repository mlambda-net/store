import { createMuiTheme } from '@material-ui/core/styles';
export const theme = createMuiTheme({
  palette: {
    type: 'light',
    primary: {
      main: "#009688",
      dark: "#00695f",
      light: "#33ab9f",
      contrastText: '#FFFFFF',
    },
    secondary: {
      main: "#f50057",
      dark: "#ab003c",
      light: "#f73378",
      contrastText: '#FFFFFF',
    },
    text: {
      primary: 'rgba(0, 0, 0, 0.87)',
      secondary: 'rgba(0, 0, 0, 0.54)',
    },
    divider: 'rgba(0, 0, 0, 0.12)',
  },
});
