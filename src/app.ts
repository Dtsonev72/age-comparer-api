import express from 'express';
import { handleProcessAge } from './controllers/ageController';

const app = express();
app.use(express.json());

app.post('/api/age', handleProcessAge);

export default app;