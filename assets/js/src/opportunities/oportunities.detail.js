import React, {Component} from 'react'
import AppBar from "@material-ui/core/AppBar";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import { withStyles } from '@material-ui/core/styles';
import SwipeableViews from 'react-swipeable-views';
import Model from "./opportunities.model";
import Metrics from "./opportunities.metrics";
import FatalAttributes from "./opportunities.fatalAttributes";

//TODO List current metrics that are associated with the opportunity. Needs to be tested
//TODO List current fatal attrabutes that are on the opportunity. Needs to be tested

const styles = theme => ({
    root: {
        backgroundColor: theme.palette.background.paper,
        width: 500,
    },
});


class OpportunityDetail extends Component {

    constructor(props) {
        super(props)

        this.state = {
            // value: null,
            metric: {},
            fatalAttribute: {}
        }
    }

    handleTabChange = (event, value) => {
        this.setState({value});
    };

    handleTabChangeIndex = index => {
        this.setState({value: index});
    };

    render() {
        const { classes, theme } = this.props;
        return (
            //TODO Add tabs for metrics and Fatal Attributes
            <div>
                <div className={classes.root}>
                    <AppBar position="static" color="default">
                        <Tabs
                            value={this.state.value}
                            onChange={this.handleTabChange}
                            indicatorColor="primary"
                            textColor="primary"
                            fullWidth
                        >
                            <Tab label="Opportunity"/>
                            <Tab label="Metrics"/>
                            <Tab label="Fatal Attributes"/>
                        </Tabs>
                    </AppBar>
                    <SwipeableViews
                        axis={theme.direction === 'rtl' ? 'x-reverse' : 'x'}
                        index={this.state.value}
                        onChangeIndex={this.handleTabChangeIndex}
                    >
                        <Model {...this.props} dir={theme.direction}/>
                        <Metrics {...this.props} dir={theme.direction}/>
                        <FatalAttributes {...this.props} dir={theme.direction}/>
                    </SwipeableViews>
                </div>
            </div>
        )
    }
}

export default withStyles(styles, { withTheme: true })(OpportunityDetail)
