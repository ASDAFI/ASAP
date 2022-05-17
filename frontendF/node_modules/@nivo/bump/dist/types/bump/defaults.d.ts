import { ModernMotionProps } from '@nivo/core';
import { BumpCommonProps, BumpPointComponent, DefaultBumpDatum } from './types';
declare const commonDefaultProps: Omit<BumpCommonProps<DefaultBumpDatum, Record<string, unknown>>, 'onMouseEnter' | 'onMouseMove' | 'onMouseLeave' | 'onClick' | 'margin' | 'theme' | 'axisRight' | 'renderWrapper'>;
export declare const bumpSvgDefaultProps: typeof commonDefaultProps & {
    pointComponent: BumpPointComponent<DefaultBumpDatum, Record<string, unknown>>;
    animate: boolean;
    motionConfig: ModernMotionProps['motionConfig'];
};
export {};
//# sourceMappingURL=defaults.d.ts.map