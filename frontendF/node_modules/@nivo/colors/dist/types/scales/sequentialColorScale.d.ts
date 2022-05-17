import { ColorInterpolatorId } from '../schemes';
export interface SequentialColorScaleConfig {
    type: 'sequential';
    scheme?: ColorInterpolatorId;
    minValue?: number;
    maxValue?: number;
}
export interface SequentialColorScaleValues {
    min: number;
    max: number;
}
export declare const sequentialColorScaleDefaults: {
    scheme: ColorInterpolatorId;
};
export declare const getSequentialColorScale: ({ scheme, minValue, maxValue, }: SequentialColorScaleConfig, values: SequentialColorScaleValues) => import("d3-scale").ScaleSequential<string, never>;
export declare const useSequentialColorScale: (config: SequentialColorScaleConfig, values: SequentialColorScaleValues) => import("d3-scale").ScaleSequential<string, never>;
//# sourceMappingURL=sequentialColorScale.d.ts.map