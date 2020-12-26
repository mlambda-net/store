import React from 'react';
import {
  AppBar,
  Button, Divider, Drawer, Grid,
  IconButton,
  List, ListItem,
  Toolbar,
  Typography
} from "@material-ui/core";
import routes from "../routes";
import MenuIcon from "@material-ui/icons/Menu";
import {theme} from "../theme";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ChevronRightIcon from "@material-ui/icons/ChevronRight";
import {Display, RouteProvider} from "@mlambda-net/core/packages/routes";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";
import clsx from "clsx";
import {ShoppingCart} from "@material-ui/icons";
import LocalizedStrings from "react-localization";



const lang = new LocalizedStrings({
  en: {
    title: 'Products',
    catalog: 'Catalog',
    menu: 'Menu'
  },
  es: {
    title: 'Productos',
    catalog: 'Catalogo',
    menu: 'Menú'
  },
});


const styles = (themes) => ({

  app: {
    height: '100%'
  },

  appBar: {
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },

  appBarShift: {
    width: `calc(100% - 240px)`,
    marginLeft: 240,
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },

  drawer: {
    width: '240px',
    flexShrink: 0,
  },

  paper: {
    width: '240px'
  },

  title: {
    flexGrow: "1",
  },

  menu: {
    display: 'flex',
    alignItems: 'center',
    margin: '8px',
    justifyContent:'space-between'
  },

  menuItem: {
    display: 'flex',
    alignItems: 'center',
  },

  menuIcon: {
    margin: "0 5px",
  },

  content: {
    height: 'calc(100vh - 64px)',
    backgroundColor: 'rgba(255,255,255,0.4)'
  }

});

class Admin extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      open: false,
    }
    this.openDrawerHandler = this.openDrawer.bind(this)
    this.closeDrawerHandler = this.closeDrawer.bind(this)
  }


  openDrawer() {
    this.setState({open: true})
  }

  closeDrawer() {
    this.setState({open: false})
  }

  render() {

    const {classes, route} = this.props;
    const {open} = this.state
    lang.setLanguage(this.props.lang);

    return (

      <RouteProvider routes={routes}>
        <div className={classes.app}>

          <AppBar position="static" className={clsx(classes.appBar, {[classes.appBarShift]: open})}>
            <Toolbar>
              <IconButton edge="start" color="inherit" onClick={this.openDrawerHandler}>
                <MenuIcon/>
              </IconButton>
              <Typography variant="h6" className={classes.title}>
                {lang.title}
              </Typography>
              <Button color="inherit">Login</Button>
            </Toolbar>
          </AppBar>
          <Drawer open={open} variant="persistent" anchor="left" className={classes.drawer}
                  classes={{paper: classes.paper}}>

            <div className={classes.menu}>
              <Typography variant="h6" color="secondary">{lang.menu}</Typography>
              <IconButton onClick={this.closeDrawerHandler}>
                {theme.direction === 'ltr' ? <ChevronLeftIcon/> : <ChevronRightIcon/>}
              </IconButton>
            </div>
            <Divider/>
            <List component="nav" aria-label="main mailbox folders">
              <ListItem button onClick={() => route.to('catalog')}>
                <div className={classes.menuItem}>
                  <ShoppingCart className={classes.menuIcon} color="secondary"/>
                  <Typography variant="subtitle2" color="secondary">{lang.catalog}</Typography>
                </div>
              </ListItem>
            </List>
          </Drawer>
          <Grid container justify="center">
            <Grid item xs={12} sm={8}>
              <div className={clsx(classes.appBar, {[classes.appBarShift]: open}, classes.content)}>
                <Display name="global"/>
              </div>
            </Grid>
          </Grid>
        </div>
      </RouteProvider>
    )
  }
}


export default withUtils(styles)(Admin);
