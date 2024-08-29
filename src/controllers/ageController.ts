import { Request, Response } from 'express';
import { processAge } from '../services/ageService';

export const handleProcessAge = (req: Request, res: Response) => {
  const ageId = processAge(req.body);
  res.status(200).json({ id: ageId });
};