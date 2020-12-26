import React from 'react';
import {ThemeProvider} from '@material-ui/core/styles';
import { theme } from './theme';
import Admin from "./pages/admin";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";
import {AuthProvider} from "@mlambda-net/core/packages/oauth/authprovider";
import settings from "./oauth";
import {Auth} from "./store/actions";
import {connect} from "react-redux";
import {LangProvider} from "@mlambda-net/core/packages/lang";


const styles = (themes) => ({

});

class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {auth:false}
    this.dispatch = props.dispatch;
    this.onLogin = this.onLogin.bind(this)
  }

  onLogin(auth) {
    this.dispatch({type: Auth.SetAuth, payload: auth})
    this.setState({auth: true})
  }


  render = () => {
    return (
      <AuthProvider settings={settings} onLogin={this.onLogin}>
        <ThemeProvider theme={theme}>
          <LangProvider lang="es">
            { this.state.auth && <Admin/>}
          </LangProvider>
        </ThemeProvider>
      </AuthProvider>
    );
  }
}

export default connect(state => state.auth)(withUtils(styles)(App));

