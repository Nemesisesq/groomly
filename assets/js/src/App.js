import React, {Component} from 'react';
import logo from './logo.svg';
import './App.css';

import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import {withStyles} from '@material-ui/core/styles'
import Button from "@material-ui/core/Button";
import {BrowserRouter as Router, Link} from "react-router-dom";
import Drawer from '@material-ui/core/Drawer'
import {RouteWithSubRoutes, routes} from "./routes";


const drawerWidth = 240;

const styles = theme => ({
    root: {
        flexGrow: 1,
    },
    appFrame: {

        zIndex: 1,
        overflow: 'hidden',
        position: 'relative',
        display: 'flex',
        width: '100%',
    },
    appBar: {
        width: `calc(100% - ${drawerWidth}px)`,
    },
    'appBar-left': {
        marginLeft: drawerWidth,
    },
    'appBar-right': {
        marginRight: drawerWidth,
    },
    drawerPaper: {
        position: 'relative',
        width: drawerWidth,
    },
    toolbar: theme.mixins.toolbar,
    content: {
        flexGrow: 1,
        backgroundColor: theme.palette.background.default,
        padding: theme.spacing.unit * 3,
    },
});



class App extends Component {
    render() {

        const { classes } = this.props;

        const drawer = (
            <Drawer
                variant="permanent"
                classes={{
                    paper: classes.drawerPaper,
                }}
            >
                <div className={classes.toolbar} />
                {routes.map((route, i) => {
                    return (
                        <Link to={route.path} key={i}>
                            <Button>{route.title}</Button>
                        </Link>
                    )
                })}


            </Drawer>
        );

        return (
            <div className="App">
                <header className="App-header">
                    <img src={logo} className="App-logo" alt="logo"/>
                    <h1 className="App-title">Welcome to Groomlyyyyy</h1>
                </header>
                <Router>
                    <div className={classes.appFrame}>
                        <AppBar
                            position="absolute"
                            className={classes.appBar}
                        >
                            <Toolbar>
                                <Typography variant="title" color="inherit" noWrap>
                                    Get Started Below
                                </Typography>
                            </Toolbar>
                        </AppBar>
                        {drawer}
                        <main className={classes.content}>
                            <div className={classes.toolbar} />
                            {routes.map((route, i) => <RouteWithSubRoutes key={i} {...route} />)}
                        </main>
                    </div>
                </Router>
            </div>
        );
    }
}

export default withStyles(styles)(App);

