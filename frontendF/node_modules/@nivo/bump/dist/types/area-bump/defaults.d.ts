import { ModernMotionProps, SvgDefsAndFill } from '@nivo/core';
import { AreaBumpCommonProps, AreaBumpComputedSerie, DefaultAreaBumpDatum } from './types';
declare const commonDefaultProps: Omit<AreaBumpCommonProps<DefaultAreaBumpDatum, Record<string, unknown>>, 'onMouseEnter' | 'onMouseMove' | 'onMouseLeave' | 'onClick' | 'margin' | 'theme' | 'renderWrapper'>;
export declare const areaBumpSvgDefaultProps: typeof commonDefaultProps & SvgDefsAndFill<AreaBumpComputedSerie<DefaultAreaBumpDatum, Record<string, unknown>>> & {
    animate: boolean;
    motionConfig: ModernMotionProps['motionConfig'];
};
export {};
//# sourceMappingURL=defaults.d.ts.map