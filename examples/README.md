# Examples

These are starter inputs for hosted AssembleFlow templates.

If you still call the workflow `batchjob`, `batchflow`, or “批处理”, those names map to the same AssembleFlow template flow.

- `text-image-v1.input.jsonl`
- `text-image-video-v1.input.jsonl`

Use them with:

```bash
./cli/assemble-flow run submit text-image-v1 -f examples/text-image-v1.input.jsonl
./cli/assemble-flow run submit text-image-video-v1 -f examples/text-image-video-v1.input.jsonl
```
