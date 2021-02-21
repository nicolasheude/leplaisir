import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import * as React from 'react';
import { Component } from 'react';

export default function SimpleSelectFormule() {
    const [age, setAge] = React.useState('');

    const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setAge(event.target.value as string);
    };

    return (
        <div>
            <FormControl className='select' 
            style={{
                margin: 25,
                minWidth: 120
                }}>
                <InputLabel id="demo-simple-select-label">Formule</InputLabel>
                <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    value={age}
                    onChange={handleChange}
                >
                    <MenuItem value={1}>Formule 1</MenuItem>
                    <MenuItem value={2}>Formule 2</MenuItem>
                    <MenuItem value={3}>Formule 3</MenuItem>
                </Select>
            </FormControl>
        </div>
    );
}

 
/*
function SimpleSelectSemaine() {
    const classes = useStyles();
    const [age, setAge] = React.useState('');

    const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setAge(event.target.value as string);
    };

    return (
        <div>
            <FormControl className='select' 
            style={{
                margin: 25,
                minWidth: 120
*/