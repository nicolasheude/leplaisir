import React from 'react';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';

export default function TextfieldsAddress() {
    return (
        <Grid container spacing={3} style={{
            margin: 10,
            minWidth: 120
            }}>
            <Grid item xs={12}>
                <TextField
                    required
                    id="address"
                    name="address"
                    label="Adresse"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
            <Grid item xs={12} sm={6}>
                <TextField
                    required
                    id="code"
                    name="code"
                    label="Code Postal"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
            <Grid item xs={12} >
                <TextField
                    required
                    id="pays"
                    name="pays"
                    label="Pays"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
            <Grid item xs={12} >
                <TextField
                    required
                    id="email"
                    name="email"
                    label="Email"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
            <Grid item xs={12} sm={6}>
                <TextField
                    required
                    id="phone"
                    name="phone"
                    label="Téléphone portable du responsable"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>

            <Grid item xs={12} sm={6}>
                <TextField
                    id="phone2"
                    name="phone2"
                    label="Second téléphone portable du responsable"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
        </Grid>
    );
}

