import React from 'react';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';

export default function TextfieldsParents() {
    return (
        <Grid container spacing={3} style={{
            margin: 10,
            minWidth: 120
            }}>
            <Grid item xs={12} sm={6}>
                <TextField
                    required
                    id="nomparent"
                    name="nomparent"
                    label="Nom du responsable légal / Représentant légal"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
            <Grid item xs={12} sm={6}>
                <TextField
                    required
                    id="prenomparent"
                    name="prenomparent"
                    label="Prénom du responsable légal / Représentant légal"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
        </Grid>
    );
}

