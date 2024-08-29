import { Age } from '../models/ageModel';

export const processAge = (ageData: Omit<typeof Age, 'id'>): string => {
    return "Empty"
};
