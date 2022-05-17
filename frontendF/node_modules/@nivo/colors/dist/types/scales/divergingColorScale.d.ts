import { ColorInterpolatorId } from '../schemes';
export interface DivergingColorScaleConfig {
    type: 'diverging';
    scheme?: ColorInterpolatorId;
    minValue?: number;
    maxValue?: number;
    divergeAt?: number;
}
export interface DivergingColorScaleValues {
    min: number;
    max: number;
}
export declare const divergingColorScaleDefaults: {
    scheme: ColorInterpolatorId;
    divergeAt: number;
};
export declare const getDivergingColorScale: ({ scheme, divergeAt, minValue, maxValue, }: DivergingColorScaleConfig, values: DivergingColorScaleValues) => import("d3-scale").ScaleDiverging<string, never>;
export declare const useDivergingColorScale: (config: DivergingColorScaleConfig, values: DivergingColorScaleValues) => import("d3-scale").ScaleDiverging<string, never>;
//# sourceMappingURL=divergingColorScale.d.ts.map