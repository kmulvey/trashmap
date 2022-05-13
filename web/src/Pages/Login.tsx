import React from "react";
import ReactDom from "react-dom";
import { DialogAuth, BoxAuth, FullWidthAuth } from "react-mui-auth-page";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import { createMuiTheme, ThemeProvider } from "@mui/material/styles";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import HomeIcon from '@mui/icons-material/HomeOutlined';

const theme = createMuiTheme({
  palette: {
    primary: {
      main: "#007fff",
    },
  },
});

export default class Login extends React.PureComponent() {
  let mode = Array.from(window.location.href.match(/\?q=(.*)/) || [])[1] || 0;
  React.useEffect(() => {}, [mode]);

  const wait = (m) => {
    console.log(m + " started");
    return new Promise((r) => {
      setTimeout(() => {
        console.log(m);
        r();
      }, 3000);
    });
  };
  const [open, setOpen] = React.useState(true);

  const handleSignIn = async ({ email, password }) => {
    await wait("SignIn");
  };
  const handleSignUp = async ({ email, name, password }) => {
    await wait("SignUp");
  };
  const handleForget = async ({ email }) => {
    await wait("Forget");
  };

  if (mode == 0)
    return (
      <Box p={5} textAlign="center">
        <Typography align="center" variant="h4" color="primary">
          <b>react-mui-auth-page</b>
        </Typography>
        {["Dialog", "Dialog Without Tabs", "FullWidth", "Box Container"].map(
          (_, key) => (
            <Box key={key} p={1}>
              <Button
                variant="contained"
                color="primary"
                onClick={() => {
                  window.location.href = `?q=${key + 1}`;
                }}
                style={{ textTransform: "none" }}
              >
                {_}
              </Button>
            </Box>
          )
        )}
      </Box>
    );
  let props = {
    open,
    onClose() {
      setOpen(false);
    },
    logoComponent: (
      <Typography variant="h5" color="primary">
        <b>My App</b>
      </Typography>
    ),
    textFieldVariant: "outlined",
    handleSignUp,
    handleForget,
    handleSignIn,
    handleSocial: {
      Google: () => {},
      Github: () => {},
      Twitter: () => {},
    },
  };
  switch (mode) {
    case "1":
      return React.createElement(DialogAuth, props);
    case "2":
      return React.createElement(DialogAuth, { ...props, hideTabs: true });
    case "3":
      return <FullWidthAuth {...props} />;
    case "4":
      return <BoxAuth {...props} />;
    default:
      return null;
  }
};

ReactDom.render(
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <div
      style={{
        position: "fixed",
        bottom: 2,
      }}
    >
      <IconButton
        aria-label="home"
        onClick={() => {
          window.location.href = "?q=0";
        }}
      >
        <HomeIcon color="action" />
      </IconButton>
      <Typography variant="caption" color="textSecondary">
        <b>Made By Arpit</b>
      </Typography>
    </div>
    <App />
  </ThemeProvider>,
  document.getElementById("root")
);
